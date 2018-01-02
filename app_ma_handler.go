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
// ma = Manage Actions
//

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"github.com/davecgh/go-spew/spew"
	"encoding/json"
	"net/url"
)

type appManageActionsData struct {
	AppPageData
	Msg string
}

func AppManageActionsHandleGet(w http.ResponseWriter, req *http.Request) {
	if ServerConfig.Features.App == false {
		http.Redirect(w, req, "/", 302)
		return
	}

	// check if id is valid
	vars := mux.Vars(req)
	id := vars["id"]

	// read data from id
	data := &appManageActionsData{
		AppPageData: AppPageData{
			PageData: PageData{
				Title: "Manage Actions",
				InApp: true,
			},
			ThingId: id,
		},
	}
	data.SetFeaturesFromConfig()
	if !data.IsIdValid() {
		AppErrorServePage(w, "An error occurred while location session data. Please try again.", id)
		return
	}
	if err := data.Deserialize(); err != nil {
		Error.Println(err)
		AppErrorServePage(w, "An error occurred while reading session data. Please try again.", id)
		return
	}

	appManageActionsServePage(w, data)

}

func AppManageActionsDataHandleGet(w http.ResponseWriter, req *http.Request) {
	if ServerConfig.Features.App == false {
		w.WriteHeader(501)
		return
	}

	// check if id is valid
	vars := mux.Vars(req)
	id := vars["id"]

	// read data from id
	data := &appManageActionsData{
		AppPageData: AppPageData{
			PageData: PageData{
				Title: "Manage Actions",
				InApp: true,
			},
			ThingId: id,
		},
	}
	data.SetFeaturesFromConfig()
	if !data.IsIdValid() {
		w.WriteHeader(500)
		fmt.Fprint(w, "Thing Id is not valid.")
		return
	}
	if err := data.Deserialize(); err != nil {
		Error.Println(err)
		w.WriteHeader(500)
		fmt.Fprint(w, "Error deserializing session data")
		return
	}
	Debug.Printf("id=%s, wtd=%s\n", id, spew.Sdump(data.wtd))

	b, err := json.Marshal(data.wtd.Actions)
	if err != nil {
		Error.Println(err)
		w.WriteHeader(500)
		fmt.Fprint(w, "Error marshaling data")
		return
	}

	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	w.Write(b)
}

func AppManageActionsHandlePost(w http.ResponseWriter, req *http.Request) {
	if ServerConfig.Features.App == false {
		w.WriteHeader(501)
		return
	}

	err := req.ParseForm()
	if err != nil {
		Debug.Printf("Error parsing create thing form: %s\n", err)
		appCreateThingServePage(w, appEntryData{Msg: "There was an error processing your data."})
	}
	mpf := req.PostForm
	Debug.Printf(spew.Sdump(mpf))

	// check if id is valid
	vars := mux.Vars(req)
	id := vars["id"]
	mafid := mpf.Get("mafid")
	Debug.Printf("got id=%s, mafid=%s\n", id, mafid)
	if id != mafid {
		AppErrorServePage(w, "An error occurred while processing form data. Please try again.", id)
		return
	}

	data := &appManageActionsData{
		AppPageData: AppPageData{
			PageData: PageData{
				Title: "Manage Actions",
				InApp: true,
			},
			ThingId: id,
		},
	}
	data.SetFeaturesFromConfig()
	if !data.IsIdValid() {
		AppErrorServePage(w, "An error occurred while location session data. Please try again.", id)
		return
	}
	if err := data.Deserialize(); err != nil {
		Error.Println(err)
		AppErrorServePage(w, "An error occurred while reading session data. Please try again.", id)
		return
	}

	parseActionsFormData(data.wtd, mpf)
	Debug.Printf("id=%s, wtd=%s\n", id, spew.Sdump(data.wtd))

	// save..
	if data.Serialize() != nil {
		Error.Println(err)
		AppErrorServePage(w, "An error occurred while writing session data. Please try again.", id)
		return
	}

	appManageActionsServePage(w, data)

}

// given the form data , this function parses all actions from it and appends these to wtd
func parseActionsFormData(wtd *WebThingDescription, formData url.Values) {
	// parse action
	wtd.NewActions()
	for idx := 1; idx < 100; idx++ {
		keyStr := fmt.Sprintf("ma_listitem_%d_val", idx)
		key := formData.Get(keyStr)
		if key == "" {
			break
		}

		keyStr = fmt.Sprintf("ma_listitem_%d_desc", idx)
		desc := formData.Get(keyStr)

		wtd.AppendAction(WebThingAction{Name: key, Description: &desc})
	}
}

func appManageActionsServePage(w http.ResponseWriter, data *appManageActionsData) {
	templates, err := NewBasicHtmlTemplateSet("app_ma.html.tpl", "app_ma_script.html.tpl")
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
