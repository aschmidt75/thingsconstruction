//    ThingsConstruction, a code generator for WoT-based models
//    Copyright (C) 2017  @aschmidt75
//
//    This program is free software: you can redistribute it and/or modify
//    it under the terms of the GNU Affero General Public License as published
//    by the Free Software Foundation, either version 3 of the License, or
//    (at your option) any later version.
//
//    This program is distributed in the hope that it will be useful,
//    but WITHOUT ANY WARRANTY; without even the implied warranty of
//    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//    GNU Affero General Public License for more details.
//
//    You should have received a copy of the GNU Affero General Public License
//    along with this program.  If not, see <http://www.gnu.org/licenses/>.
//
package main

// app_ct_handler.go
//
// app_ct_handler is responsible for delivering the starting page of the app, where
// users can create a new thing with a basic name and description.
// It generates a UUID and saves the basic definition in a JSON file, redirecting users
// to the next step of adding properties etc. to the Thing Description.

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/satori/go.uuid"
	"net/http"
	"strings"
	"github.com/davecgh/go-spew/spew"
)

type appEntryData struct {
	AppPageData
	CtfName string
	CtfDesc string
	CtfType string
}

func AppCreateThingHandleGet(w http.ResponseWriter, req *http.Request) {
	if ServerConfig.Features.App == false {
		http.Redirect(w, req, "/", 302)
		return
	}

	data := &appEntryData{
		AppPageData: AppPageData{
			PageData: PageData{
				Title: "Create Thing Description",
			},
		},
	}
	data.SetFeaturesFromConfig()
	data.InApp = true

	vars := mux.Vars(req)
	data.AppPageData.ThingId = vars["id"]
	if data.AppPageData.ThingId != "" {
		Debug.Printf("Loading data for id=%s\n", data.AppPageData.ThingId)

		if err := data.Deserialize(); err != nil {
			Error.Printf("Unable to load data, err=%s\n", err)
		} else {
			data.CtfName = data.AppPageData.wtd.Name
			data.CtfDesc = *data.AppPageData.wtd.Description
			data.CtfType = data.AppPageData.wtd.Type
		}
	}

	appCreateThingServePage(w, *data)
}

func AppCreateThingHandlePost(w http.ResponseWriter, req *http.Request) {
	if ServerConfig.Features.App == false {
		http.Redirect(w, req, "/", 302)
		return
	}

	err := req.ParseForm()
	if err != nil {
		Debug.Printf("Error parsing create thing form: %s\n", err)
		appCreateThingServePage(w, appEntryData{
			AppPageData:AppPageData{
				Message: "There was an error processing your data.",
			},
		})
		return
	}
	ctf := req.PostForm

	// check if id is valid
	vars := mux.Vars(req)
	id := vars["id"]
	if id != "" {
		// this is an edit of an existing ThingId
		data := &appManageActionsData{
			AppPageData: AppPageData{
				PageData: PageData{
					Title: "Manage Actions",
					InApp: true,
				},
				ThingId: id,
			},
		}
		if !data.IsIdValid() {
			appCreateThingServePage(w, appEntryData{
				AppPageData:AppPageData{
					Message: "There was an error locating WoT data by ID.",
				},
			})
			return
		}
		if err := data.Deserialize(); err != nil {
			Error.Println(err)
			appCreateThingServePage(w, appEntryData{
				AppPageData:AppPageData{
					Message: "There was an error locating WoT data by ID.",
				},
			})
		}

		data.wtd.Name = ctf.Get("ctf_name")
		data.wtd.Type = ctf.Get("ctf_type")
		data.wtd.Description = new(string)
		*data.wtd.Description = ctf.Get("ctf_desc")
		Debug.Printf("id=%s, wtd=%s\n", id, spew.Sdump(data.wtd))

		// save..
		if data.Serialize() != nil {
			Error.Println(err)
			appCreateThingServePage(w, appEntryData{
				AppPageData:AppPageData{
					Message: "There was an error writing session data.",
				},
			})
			return
		}
		// redirect to next steps
		http.Redirect(w, req, "/app/"+id+"/framework", http.StatusFound)

	} else {
		// create new one
		data := &appEntryData{
			AppPageData: AppPageData{
				PageData: PageData{
					Title: "Create new Thing Description",
					InApp: true,
				},
			},
			CtfName: ctf.Get("ctf_name"),
			CtfDesc: strings.TrimSpace(ctf.Get("ctf_desc")),
			CtfType: ctf.Get("ctf_type"),
		}
		data.SetFeaturesFromConfig()

		id, err := appEntryCreateThing(data)
		if err != nil {
			data.AppPageData.Message = "There was an error creating your Thing Description. Please try again."
			appCreateThingServePage(w, *data)
		} else {
			// redirect to next steps
			http.Redirect(w, req, "/app/"+id+"/framework", http.StatusFound)
		}
	}

}

// creates a new Thing Description json, puts basic data in it,
// returns unique thing id
func appEntryCreateThing(data *appEntryData) (string, error) {

	data.AppPageData.ThingId = uuid.NewV4().String()

	data.AppPageData.wtd = &WebThingDescription{
		Name:        data.CtfName,
		Description: new(string),
		Type:        data.CtfType,
	}
	*data.AppPageData.wtd.Description = data.CtfDesc

	err := data.AppPageData.Serialize()
	if err != nil {
		return "", err
	}

	return data.AppPageData.ThingId, nil
}

func appCreateThingServePage(w http.ResponseWriter, data appEntryData) {
	templates, err := NewBasicHtmlTemplateSet("app_ct.html.tpl", "app_ct_script.html.tpl")
	if err != nil {
		Error.Fatalf("Fatal error creating template set: %s\n", err)
		panic(err)
	}

	Verbose.Printf("appCreateThingServePage: %#v\n", data)

	err = templates.ExecuteTemplate(w, "root", data)
	if err != nil {
		Error.Printf("Error executing template: %s\n", err)
		w.WriteHeader(500)
		fmt.Fprint(w, "There was an internal error.")
	}

}
