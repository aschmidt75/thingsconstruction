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
)

type appManagePropertiesData struct {
	AppPageData
}

func AppManagePropertiesHandleGet(w http.ResponseWriter, req *http.Request) {
	if ServerConfig.Features.App == false {
		http.Redirect(w, req, "/", 302)
		return
	}

	// user selected a generator.
	// check if id is valid
	vars := mux.Vars(req)
	id := vars["id"]

	// read data from id

	data := &appManagePropertiesData{
		AppPageData: AppPageData{
			PageData: PageData{
				Title: "THNGS:CONSTR - Manage Properties",
				InApp: true,
			},
			ThingId: id,
		},
	}
	data.SetFeaturesFromConfig()

	appManagePropertiesServePage(w, data)

}

func appManagePropertiesServePage(w http.ResponseWriter, data *appManagePropertiesData) {
	templates, err := NewBasicHtmlTemplateSet("app_mp.html.tpl.tpl", "app_mp_script.html.tpl")
	if err != nil {
		Error.Fatalf("Fatal error creating template set: %s\n", err)
	}

	Verbose.Printf("appManagePropertiesServePage: %#v\n", data)

	err = templates.ExecuteTemplate(w, "root", data)
	if err != nil {
		Error.Printf("Error executing template: %s\n", err)
		w.WriteHeader(500)
		fmt.Fprint(w, "There was an internal error.")
	}

}
