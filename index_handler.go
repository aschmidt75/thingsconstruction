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
//    This program is dual-licensed. For commercial licensing options, please
//    contact the author(s).
//
package main

import (
	"fmt"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	templates, err := NewBasicHtmlTemplateSet("index.html.tpl", "index_script.html.tpl")
	if err != nil {
		Error.Fatalf("Fatal error creating template set: %s\n", err)
	}

	data := PageData{
		Title: "Index",
	}
	data.SetFeaturesFromConfig()

	err = templates.ExecuteTemplate(w, "root", data)
	if err != nil {
		Error.Printf("Error executing template: %s", err)
		w.WriteHeader(500)
		fmt.Fprint(w, "There was an internal error.")
	}
}
