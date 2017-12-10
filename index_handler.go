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
	"net/http"
)

func IndexHandler(w http.ResponseWriter, req *http.Request) {
	templates, err := NewBasicHtmlTemplateSet("index.html.tpl")
	if err != nil {
		Error.Fatalf( "Fatal error creating template set: %s\n", err)
	}

	context := PageData{
		Title: "Index",
	}
	err = templates.ExecuteTemplate(w, "root", context)
	if err != nil {
		Error.Printf("Error executing template: %s", err)

		// TODO: Send error page
	}
}