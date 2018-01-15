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
// mp = manage properties
//

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"github.com/davecgh/go-spew/spew"
	"strings"
	"encoding/json"
	"strconv"
	"net/url"
)

type appManagePropertiesData struct {
	AppPageData
	Msg string
}

func appManagePropertiesNewPageData(id string) (*appManagePropertiesData) {
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
		return nil
	}
	if err := data.Deserialize(); err != nil {
		Error.Println(err)
		return nil
	}
	data.SetTocInfo()

	return data
}

func AppManagePropertiesHandleGet(w http.ResponseWriter, req *http.Request) {
	if ServerConfig.Features.App == false {
		http.Redirect(w, req, "/", 302)
		return
	}

	// check if id is valid
	vars := mux.Vars(req)
	id := vars["id"]

	data := appManagePropertiesNewPageData(id)
	if data == nil {
		AppErrorServePage(w, "An error occurred while reading session data. Please try again.", id)
		return
	}

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

	data := appManagePropertiesNewPageData(id)
	if data == nil {
		w.WriteHeader(500)
		fmt.Fprint(w, "Thing Id is not valid or Error deserializing session data.")
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
		w.WriteHeader(http.StatusNotImplemented)
		return
	}

	err := req.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "There was an error processing your data. Please try again.")
		Debug.Printf("Error parsing create thing form: %s\n", err)
	}
	mpf := req.PostForm
	Debug.Printf(spew.Sdump(mpf))

	// check if id is valid
	vars := mux.Vars(req)
	id := vars["id"]
	mpfid := mpf.Get("mpfid")
	Debug.Printf("got id=%s, mpfid=%s\n", id, mpfid)
	if id != mpfid {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "An error occurred while processing form data. Please try again.")
		return
	}

	data := appManagePropertiesNewPageData(id)
	if data == nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "An error occurred while location/reading session data. Please try again.")
		return
	}

	parsePropertiesFormData(data.wtd, mpf)
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

// given the form data (mpf), this function parses all properties
// from form and appends these to wtd
func parsePropertiesFormData(wtd *WebThingDescription, mpf url.Values) {
	// parse properties
	wtd.NewProperties()
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

		switch keyArr[1] {
		case "b":
			{
				wtd.AppendProperty(WebThingProperty{Name: keyArr[0], Type: "Boolean", Description: &desc})
			}
		case "s":
			{
				var maxLength *int = nil
				if len(keyArr) >= 3 {
					if x, err := strconv.Atoi(keyArr[2]); err == nil {
						maxLength = new(int)
						*maxLength = x
					}
				}
				wtd.AppendProperty(WebThingProperty{Name: keyArr[0], Type: "String", MaxLength: maxLength, Description: &desc})
			}
		case "i":
			{
				var minVal *float64 = nil
				var maxVal *float64 = nil
				if len(keyArr) >= 3 {
					if x, err := strconv.Atoi(keyArr[2]); err == nil {
						minVal = new(float64)
						*minVal = float64(x)
					}
				}
				if len(keyArr) >= 4 {
					if x, err := strconv.Atoi(keyArr[3]); err == nil {
						maxVal = new(float64)
						*maxVal = float64(x)
					}
				}
				wtd.AppendProperty(WebThingProperty{Name: keyArr[0], Type: "Integer", Min: minVal, Max: maxVal, Description: &desc})
			}
		case "f":
			{
				var minVal *float64 = nil
				var maxVal *float64 = nil
				if len(keyArr) >= 3 {
					if x, err := strconv.ParseFloat(keyArr[2], 64); err == nil {
						minVal = new(float64)
						*minVal = x
					}
				}
				if len(keyArr) >= 4 {
					if x, err := strconv.ParseFloat(keyArr[3], 64); err == nil {
						maxVal = new(float64)
						*maxVal = x
					}
				}
				wtd.AppendProperty(WebThingProperty{Name: keyArr[0], Type: "Float", Min: minVal, Max: maxVal, Description: &desc})
			}
		}

	}
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
