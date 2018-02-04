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
	"html/template"
	"io/ioutil"
	"net/http"
	"path/filepath"
)

var whiteListPages = map[string]string{
	"about":      "about.html",
	"imprint":    "imprint.html",
	"imprint_de": "imprint_de.html",
	"privacy":    "privacy.html",
	"privacy_de": "privacy_de.html",
}

type staticPageData struct {
	PageData
	HtmlOutput template.HTML
}

func ServeNotFound(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(404)

	bp := filepath.Join(ServerConfig.Paths.StaticPagesPath, "notfound.html.tpl")
	Debug.Printf("serving not found page %s\n", bp)

	tplBytes, err := ioutil.ReadFile(bp)
	if err != nil {
		Debug.Printf("ServeNotFound: err=%s\n", err)
		return
	}

	staticPagesServePage(w, staticPageData{
		PageData: PageData{
			Title: "Page not found", // TODO
		},
		HtmlOutput: template.HTML(tplBytes),
	})
}

func StaticPageHandler(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	pageName := vars["page"]

	// look up page by name
	bp, ok := whiteListPages[pageName]
	if ok {
		bp = filepath.Join(ServerConfig.Paths.StaticPagesPath, bp)
		Debug.Printf("serving static page %s\n", bp)

		tplBytes, err := ioutil.ReadFile(bp)
		if err != nil {
			Error.Printf("Error reading page by name %s\n", pageName)
			ServeNotFound(w, req)
			return
		}

		staticPagesServePage(w, staticPageData{
			PageData: PageData{
				Title: pageName, // TODO
			},
			HtmlOutput: template.HTML(tplBytes),
		})
	} else {
		ServeNotFound(w, req)
	}

}

var StaticPagesTemplates *template.Template

// initializes template set. Static Pages are
// staticpage.html.tpl, staticpage_script.html.tpl
// plus the inner content parts
func initializeTemplates() {
	if StaticPagesTemplates == nil {
		Debug.Printf("Initializing templates for static pages")
		var err error
		StaticPagesTemplates, err = NewBasicHtmlTemplateSet("staticpage.html.tpl", "staticpage_script.html.tpl")
		if err != nil {
			Error.Fatalf("Fatal error creating template set: %s\n", err)
		}
	}
}

func staticPagesServePage(w http.ResponseWriter, data staticPageData) {
	initializeTemplates()
	data.SetFeaturesFromConfig()

	err := StaticPagesTemplates.ExecuteTemplate(w, "root", data)
	if err != nil {
		Error.Printf("Error executing template: %s\n", err)
		w.WriteHeader(500)
		fmt.Fprint(w, "There was an internal error.")
	}

}
