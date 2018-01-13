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

//
// ma = Manage Events
//

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"github.com/davecgh/go-spew/spew"
	"encoding/json"
	"net/url"
)

type appManageEventsData struct {
	AppPageData
	Msg string
}

func appManageEventsNewPageData(id string) (*appManageEventsData) {
	// read data from id
	data := &appManageEventsData{
		AppPageData: AppPageData{
			PageData: PageData{
				Title: "Manage Events",
				InApp: true,
			},
			ThingId: id,
		},
	}
	data.SetFeaturesFromConfig()
	if !data.IsIdValid() {
		return nil
	}
	if err := data.Deserialize(); err != nil {
		Error.Println(err)
		return nil
	}
	data.SetTocInfo()

	return data
}

func AppManageEventsHandleGet(w http.ResponseWriter, req *http.Request) {
	if ServerConfig.Features.App == false {
		http.Redirect(w, req, "/", 302)
		return
	}

	// check if id is valid
	vars := mux.Vars(req)
	data := appManageEventsNewPageData(vars["id"])
	if data == nil {
		AppErrorServePage(w, "An error occurred while reading session data. Please try again.", vars["id"])
	}

	appManageEventsServePage(w, data)

}

func AppManageEventsDataHandleGet(w http.ResponseWriter, req *http.Request) {
	if ServerConfig.Features.App == false {
		w.WriteHeader(501)
		return
	}

	// check if id is valid
	vars := mux.Vars(req)
	id := vars["id"]

	data := appManageEventsNewPageData(id)
	if data == nil {
		w.WriteHeader(500)
		fmt.Fprint(w, "Error deserializing session data")
		return
	}
	Debug.Printf("id=%s, wtd=%s\n", id, spew.Sdump(data.wtd))

	b, err := json.Marshal(data.wtd.Events)
	if err != nil {
		Error.Println(err)
		w.WriteHeader(500)
		fmt.Fprint(w, "Error marshaling data")
		return
	}
	Debug.Printf("events-data: %s\n", b)

	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	w.Write(b)
}

func AppManageEventsHandlePost(w http.ResponseWriter, req *http.Request) {
	if ServerConfig.Features.App == false {
		w.WriteHeader(http.StatusNotImplemented)
		return
	}

	err := req.ParseForm()
	if err != nil {
		Debug.Printf("Error parsing events form: %s\n", err)
		appCreateThingServePage(w, appEntryData{
			AppPageData: AppPageData{
				Message: "There was an error processing your data.",
			}})
	}
	formData := req.PostForm
	Debug.Printf(spew.Sdump(formData))

	// check if id is valid
	vars := mux.Vars(req)
	id := vars["id"]
	mefid := formData.Get("mefid")
	Debug.Printf("got id=%s, mefid=%s\n", id, mefid)
	if id != mefid {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "An error occurred while processing form data. Please try again.")
		return
	}

	data := &appManageEventsData{
		AppPageData: AppPageData{
			PageData: PageData{
				Title: "Manage Events",
				InApp: true,
			},
			ThingId: id,
		},
	}
	data.SetFeaturesFromConfig()
	if !data.IsIdValid() {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "An error occurred while location session data. Please try again.")
		return
	}
	if err := data.Deserialize(); err != nil {
		Error.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "An error occurred while reading session data. Please try again.")
		return
	}

	parseEventsFormData(data.wtd, formData)
	Debug.Printf("id=%s, wtd=%s\n", id, spew.Sdump(data.wtd))

	// save..
	if data.Serialize() != nil {
		Error.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "An error occurred while writing session data. Please try again.")
		return
	}

	w.WriteHeader(http.StatusOK)
}

// given the form data , this function parses all events from it and appends these to wtd
func parseEventsFormData(wtd *WebThingDescription, formData url.Values) {
	// parse Event
	wtd.NewEvents()
	for idx := 1; idx < 100; idx++ {
		keyStr := fmt.Sprintf("me_listitem_%d_val", idx)
		key := formData.Get(keyStr)
		if key == "" {
			break
		}

		keyStr = fmt.Sprintf("me_listitem_%d_desc", idx)
		desc := formData.Get(keyStr)

		wtd.AppendEvent(WebThingEvent{Name: key, Description: &desc})
	}
}

func appManageEventsServePage(w http.ResponseWriter, data *appManageEventsData) {
	templates, err := NewBasicHtmlTemplateSet("app_me.html.tpl", "app_me_script.html.tpl")
	if err != nil {
		Error.Fatalf("Fatal error creating template set: %s\n", err)
	}

	err = templates.ExecuteTemplate(w, "root", data)
	if err != nil {
		Error.Printf("Error executing template: %s\n", err)
		w.WriteHeader(500)
		fmt.Fprint(w, "There was an internal error.")
	}

}
