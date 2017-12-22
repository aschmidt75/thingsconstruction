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
	"github.com/gorilla/mux"
	"github.com/satori/go.uuid"
	"net/http"
	"strings"
)

type appEntryData struct {
	PageData
	Msg     string
	CtfName string
	CtfDesc string
	CtfType string
}

func AppCreateThingHandleGet(w http.ResponseWriter, req *http.Request) {
	if ServerConfig.Features.App == false {
		http.Redirect(w, req, "/", 302)
		return
	}

	data := &appEntryData{}
	data.SetFeaturesFromConfig()
	data.InApp = true

	vars := mux.Vars(req)
	id := vars["id"]
	if id != "" {
		Debug.Printf("Loading data for id=%s\n", id)
		wtd := WebThingDescription{}

		if err := wtd.Deserialize(id); err != nil {
			Error.Printf("Unable to load data: %s\n", err)
		} else {
			data.CtfName = wtd.Name
			data.CtfDesc = wtd.Description
			data.CtfType = wtd.Type
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
		appCreateThingServePage(w, appEntryData{Msg: "There was an error processing your data."})
	}
	ctf := req.PostForm

	data := &appEntryData{
		PageData: PageData{
			Title: "THNGS:CONSTR - Create new Thing Description",
			InApp: true,
		},
		CtfName: ctf.Get("ctf_name"),
		CtfDesc: strings.TrimSpace(ctf.Get("ctf_desc")),
		CtfType: ctf.Get("ctf_type"),
	}
	data.SetFeaturesFromConfig()

	// validate
	if data.CtfType != "thing" {
		data.Msg = "I'm sorry, at present only type \"Thing\" is supported."
	}

	id, err := appEntryCreateThing(data)
	if err != nil {
		data.Msg = "There was an error creating your Thing Description. Please try again."
		appCreateThingServePage(w, *data)
	} else {
		// redirect to next steps
		http.Redirect(w, req, "/app/"+id+"/framework", http.StatusFound)
		appCreateThingServePage(w, *data)

	}
}

// creates a new Thing Description json, puts basic data in it,
// returns unique id
func appEntryCreateThing(data *appEntryData) (string, error) {

	id := uuid.NewV4().String()

	wtd := WebThingDescription{
		Name:        data.CtfName,
		Description: data.CtfDesc,
		Type:        data.CtfType,
	}

	err := wtd.Serialize(id)
	if err != nil {
		return "", err
	}

	return id, nil
}

func appCreateThingServePage(w http.ResponseWriter, data appEntryData) {
	templates, err := NewBasicHtmlTemplateSet("app_ct.html.tpl", "app_ct_script.html.tpl")
	if err != nil {
		Verbose.Printf("Fatal error creating template set: %s\n", err)
		panic(err)
	}

	Verbose.Printf("appCreateThingServePage: %#v\n", data)

	err = templates.ExecuteTemplate(w, "root", data)
	if err != nil {
		Verbose.Printf("Error executing template: %s\n", err)
	}

}
