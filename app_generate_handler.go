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

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"github.com/davecgh/go-spew/spew"
	"encoding/json"
	"github.com/satori/go.uuid"
)

type appGenerateData struct {
	AppPageData
	Msg string
	Accepted bool
}

func appGenerateNewPageData(id string) (*appGenerateData) {
	// read data from id
	data := &appGenerateData{
		AppPageData: AppPageData{
			PageData: PageData{
				Title: "Generate WoT code",
				InApp: true,
			},
			ThingId: id,
		},
		Accepted: false,
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

func AppGenerateHandleGet(w http.ResponseWriter, req *http.Request) {
	if ServerConfig.Features.App == false {
		http.Redirect(w, req, "/", 302)
		return
	}

	// check if id is valid
	vars := mux.Vars(req)
	data := appGenerateNewPageData(vars["id"])
	if data == nil {
		AppErrorServePage(w, "An error occurred while reading session data. Please try again.", vars["id"])
	}

	appGenerateServePage(w, data)

}


func AppGenerateDataHandleGet(w http.ResponseWriter, req *http.Request) {
	if ServerConfig.Features.App == false {
		w.WriteHeader(501)
		return
	}

	// check if id is valid
	vars := mux.Vars(req)
	id := vars["id"]

	data := appGenerateNewPageData(id)
	if data == nil {
		w.WriteHeader(500)
		fmt.Fprint(w, "Error deserializing session data")
		return
	}
	Debug.Printf("id=%s, wtd=%s\n", id, spew.Sdump(data.wtd))

	t, err := ReadGeneratorsConfig()
	if err != nil {
		Error.Printf("Unable to present generators. FIX CONFIG!\n")
		w.WriteHeader(500)
		fmt.Fprint(w, "Error loading generator data")
		return
	}

	pageData := struct{
		Wtd *WebThingDescription `json:"wtd"`
		Target *AppGenTarget `json:"target"`
	}{ Wtd: data.wtd, Target: t.AppGenTargetById(data.md.SelectedGeneratorId)}

	b, err := json.Marshal(pageData)
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

func AppGenerateWtdHandleGet(w http.ResponseWriter, req *http.Request) {
	if ServerConfig.Features.App == false {
		w.WriteHeader(501)
		return
	}

	// check if id is valid
	vars := mux.Vars(req)
	id := vars["id"]

	data := appGenerateNewPageData(id)
	if data == nil {
		w.WriteHeader(500)
		fmt.Fprint(w, "Error deserializing session data")
		return
	}

	b, err := json.MarshalIndent(data.wtd, "", "\t")
	if err != nil {
		Error.Println(err)
		w.WriteHeader(500)
		fmt.Fprint(w, "Error marshaling data")
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Content-Disposition", "attachment; filename=\""+data.ThingId+".json\"")
	w.WriteHeader(200)
	w.Write(b)
}

func AppGenerateAcceptHandlePost(w http.ResponseWriter, req *http.Request) {
	if ServerConfig.Features.App == false {
		w.WriteHeader(http.StatusNotImplemented)
		return
	}

	err := req.ParseForm()
	if err != nil {
		Debug.Printf("Error parsing generate form: %s\n", err)
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

	pageData := struct{
		Id string
		Token string
	}{ Id: id, Token: uuid.NewV4().String()}

	b, err := json.Marshal(pageData)
	if err != nil {
		Error.Println(err)
		w.WriteHeader(500)
		fmt.Fprint(w, "Error marshaling data")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	w.Write(b)
}

/*
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
*/

func appGenerateServePage(w http.ResponseWriter, data *appGenerateData) {
	templates, err := NewBasicHtmlTemplateSet("app_generate.html.tpl", "app_generate_script.html.tpl")
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
