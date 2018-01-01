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
	"strings"
	"encoding/json"
	"strconv"
)

type appManagePropertiesData struct {
	AppPageData
	Msg string

}

func AppManagePropertiesHandleGet(w http.ResponseWriter, req *http.Request) {
	if ServerConfig.Features.App == false {
		http.Redirect(w, req, "/", 302)
		return
	}

	// check if id is valid
	vars := mux.Vars(req)
	id := vars["id"]

	// read data from id

	data := &appManagePropertiesData{
		AppPageData: AppPageData{
			PageData: PageData{
				Title: "Manage Properties",
				InApp: true,
			},
			ThingId: id,
		},
	}
	data.SetFeaturesFromConfig()

	appManagePropertiesServePage(w, data)

}

func AppManagePropertiesDataHandleGet(w http.ResponseWriter, req *http.Request) {
	if ServerConfig.Features.App == false {
		http.Redirect(w, req, "/", 302)
		return
	}

	// check if id is valid
	vars := mux.Vars(req)
	id := vars["id"]

	// read data from id
	data := &appManagePropertiesData{
		AppPageData: AppPageData{
			PageData: PageData{
				Title: "Manage Properties",
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

	b, err := json.Marshal(data.wtd.Properties)
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

func AppManagePropertiesHandlePost(w http.ResponseWriter, req *http.Request) {
	if ServerConfig.Features.App == false {
		http.Redirect(w, req, "/", 302)
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
	mpfid := mpf.Get("mpfid")
	Debug.Printf("got id=%s, mpfid=%s\n", id, mpfid)
	if id != mpfid {
		AppErrorServePage(w, "An error occurred while processing form data. Please try again.", id)
		return
	}

	data := &appManagePropertiesData{
		AppPageData: AppPageData{
			PageData: PageData{
				Title: "Manage Properties",
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
	Debug.Printf("id=%s, wtd=%s\n", id, spew.Sdump(data.wtd))

	// parse properties
	data.wtd.NewProperties()
	for idx := 1; idx < 100; idx++ {
		keyStr := fmt.Sprintf("mp_listitem_%d_val", idx)
		key := mpf.Get(keyStr)
		if key == "" {
			break
		}

		keyArr := strings.Split(key, ";")
		Debug.Printf("idx=%d, key=%#v\n", idx, keyArr)
		if len(keyArr) < 2 {
			// error, we need at least name and type.
			Verbose.Printf("Invalid wtd property spec: %s\n", key)
			continue
		}

		keyStr = fmt.Sprintf("mp_listitem_%d_desc", idx)
		desc := mpf.Get(keyStr)
		Debug.Printf("idx=%d, desc=%s\n", idx, desc)

		switch keyArr[1] {
		case "b": {
			data.wtd.AppendProperty(WebThingProperty{Name: keyArr[0], Type: "Boolean", Description: &desc})
		}
		case "s": {
			var maxLength *int = nil
			if len(keyArr) >= 3 {
				if x, err := strconv.Atoi(keyArr[2]); err == nil {
					maxLength = new(int)
					*maxLength = x
				}
			}
			data.wtd.AppendProperty(WebThingProperty{Name: keyArr[0], Type: "String", MaxLength: maxLength, Description: &desc})
		}
		case "i": {
			var minVal *int = nil
			var maxVal *int = nil
			if len(keyArr) >= 3 {
				if x, err := strconv.Atoi(keyArr[2]); err == nil {
					minVal = new(int)
					*minVal = x
				}
			}
			if len(keyArr) >= 4 {
				if x, err := strconv.Atoi(keyArr[2]); err == nil {
					maxVal = new(int)
					*maxVal = x
				}
			}
			data.wtd.AppendProperty(WebThingProperty{Name: keyArr[0], Type: "Integer", Min: minVal, Max: maxVal, Description: &desc})
		}
		case "f": {
			var minVal *int = nil
			var maxVal *int = nil
			if len(keyArr) >= 3 {
				if x, err := strconv.Atoi(keyArr[2]); err == nil {
					minVal = new(int)
					*minVal = x
				}
			}
			if len(keyArr) >= 4 {
				if x, err := strconv.Atoi(keyArr[2]); err == nil {
					maxVal = new(int)
					*maxVal = x
				}
			}
			data.wtd.AppendProperty(WebThingProperty{Name: keyArr[0], Type: "Float", Min: minVal, Max: maxVal, Description: &desc})
		}
		}

	}

	Debug.Printf("id=%s, wtd=%s\n", id, spew.Sdump(data.wtd))
	// save..
	if data.Serialize() != nil {
		Error.Println(err)
		AppErrorServePage(w, "An error occurred while writing session data. Please try again.", id)
		return
	}

	appManagePropertiesServePage(w, data)

}

func appManagePropertiesServePage(w http.ResponseWriter, data *appManagePropertiesData) {
	templates, err := NewBasicHtmlTemplateSet("app_mp.html.tpl", "app_mp_script.html.tpl")
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
