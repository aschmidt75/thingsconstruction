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
	"github.com/davecgh/go-spew/spew"
	"github.com/gorilla/mux"
	"net/http"
)

type appGenParamsData struct {
	AppPageData
	NumGenerators int
	AppGenTargets *AppGenTargets
}

func appGenParamsNewPageData(id string) (*appGenParamsData, error) {

	t, err := ReadGeneratorsConfig()
	if err != nil {
		Error.Printf("Unable to present generators. FIX CONFIG!\n")
		return nil, err
	}

	var data = &appGenParamsData{
		AppPageData: AppPageData{
			PageData: PageData{
				Title: "Choose Embedded Development Framework",
			},
			ThingId: id,
		},
		NumGenerators: 1,
		AppGenTargets: t,
	}
	data.SetFeaturesFromConfig()
	data.InApp = true
	if !data.IsIdValid() {
		Error.Printf("Invalid ID=%s\n", id)
		return nil, &AppError{"Unable to locate WoT data for given ID"}
	}
	if err := data.Deserialize(); err != nil {
		Error.Println(err)
		return nil, &AppError{"Unable to locate WoT data for given ID"}
	}

	data.SetTocInfo()

	return data, nil
}

func AppChooseFrameworkHandleGet(w http.ResponseWriter, req *http.Request) {
	if ServerConfig.Features.App == false {
		http.Redirect(w, req, "/", 302)
		return
	}

	vars := mux.Vars(req)
	id := vars["id"]

	data, err := appGenParamsNewPageData(id)
	if err != nil {
		// send back to create page
		http.Redirect(w, req, "/app", 302)
		return
	}
	appGenParamsServePage(w, *data)
}

func AppChooseFrameworkHandlePost(w http.ResponseWriter, req *http.Request) {
	if ServerConfig.Features.App == false {
		http.Redirect(w, req, "/", 302)
		return
	}

	err := req.ParseForm()
	if err != nil {
		Debug.Printf("Error parsing choose framework form: %s\n", err)
		appCreateThingServePage(w, appEntryData{
			AppPageData: AppPageData{
				Message: "There was an error processing your data.",
			}})
	}
	ctf := req.PostForm
	Debug.Printf(spew.Sdump(ctf))

	// user selected a generator.
	// check if id is valid
	vars := mux.Vars(req)
	id := vars["id"]
	cfid := ctf.Get("cfid")
	cfs := ctf.Get("cfs")

	Debug.Printf("got id=%s, cfs=%s, cfid=%s\n", id, cfs, cfid)
	if id != cfid {
		AppErrorServePage(w, "An error occurred while processing form data. Please try again.", id)
		return
	}

	data := &AppPageData{
		ThingId: id,
	}
	if !data.IsIdValid() {
		AppErrorServePage(w, "An error occurred while location session data. Please try again.", id)
		return
	}
	if err := data.Deserialize(); err != nil {
		Error.Println(err)

		AppErrorServePage(w, "An error occurred while reading session data. Please try again.", id)
		return
	}
	Debug.Printf("id=%s, wtd=%s\n", id, spew.Sdump(data.wtd))

	// write generator to meta data file
	data.md = &GeneratorMetaData{
		SelectedGeneratorId: cfs,
	}
	if err := data.Serialize(); err != nil {
		Error.Println(err)

		AppErrorServePage(w, "An error occurred while storing session data. Please try again.", id)
		return
	}

	// redirect to manage properties
	http.Redirect(w, req, "/app/"+id+"/properties", 302)
}

func appGenParamsServePage(w http.ResponseWriter, data appGenParamsData) {
	templates, err := NewBasicHtmlTemplateSet("app_cf.html.tpl", "app_cf_script.html.tpl")
	if err != nil {
		Error.Fatalf("Fatal error creating template set: %s\n", err)
	}

	if err = templates.ExecuteTemplate(w, "root", data); err != nil {
		Error.Printf("Error executing template: %s\n", err)
	}

}
