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
	"github.com/gorilla/mux"
	"html/template"
	"io/ioutil"
	"net/http"
)

var whiteListPages = map[string]string{
	"about": "static-content/about.html.tpl",
}

type staticPageData struct {
	PageData
	HtmlOutput template.HTML
}

func StaticPageHandler(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	pageName := vars["page"]

	// look up page by name
	bp, ok := whiteListPages[pageName]
	if ok {
		tplBytes, err := ioutil.ReadFile(bp)
		if err != nil {
			Error.Printf("Error reading page by name %s\n", pageName)
			w.WriteHeader(404)
			return
		}

		staticPagesServePage(w, staticPageData{
			PageData: PageData{
				Title: pageName, // TODO
			},
			HtmlOutput: template.HTML(tplBytes),
		})
	} else {
		w.WriteHeader(404)
	}

}

func staticPagesServePage(w http.ResponseWriter, data staticPageData) {
	templates, err := NewBasicHtmlTemplateSet("staticpage.html.tpl", "staticpage_script.html.tpl")
	if err != nil {
		Error.Printf("Fatal error creating template set: %s\n", err)
		panic(err)
	}

	err = templates.ExecuteTemplate(w, "root", data)
	if err != nil {
		Error.Printf("Error executing template: %s\n", err)
	}

}
