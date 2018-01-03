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
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

type GeneratorMetaData struct {
	SelectedGeneratorId string `json:"genid"`
}

func (gmd *GeneratorMetaData) Serialize(id string, fileName string) error {
	b, err := json.Marshal(gmd)
	if err != nil {
		return err
	}

	Debug.Printf("%s %s", fileName, b)
	return ioutil.WriteFile(fileName, b, 0640)
}

func (gmd *GeneratorMetaData) Deserialize(id string, fileName string) error {
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		Debug.Printf("Error reading %s: %s", fileName, err)
		return err
	}

	err = json.Unmarshal(b, gmd)
	if err != nil {
		Debug.Printf("Error parsing %s: %s", fileName, err)
		Debug.Println(string(b))
	}
	return err
}

type AppPageData struct {
	PageData
	ThingId string
	TocInfo map[string]string
	wtd     *WebThingDescription
	md      *GeneratorMetaData
}

func (ap *AppPageData) SetTocInfo() {
	ap.TocInfo = map[string]string{}

	if ap.wtd != nil {
		var l int
		l = 0
		if ap.wtd.Properties != nil {
			l = len(*ap.wtd.Properties)
		}
		ap.TocInfo["num_properties"] = fmt.Sprintf("%d", l)

		l = 0
		if ap.wtd.Actions != nil {
			l = len(*ap.wtd.Actions)
		}
		ap.TocInfo["num_actions"] = fmt.Sprintf("%d", l)

		l = 0
		if ap.wtd.Events != nil {
			l = len(*ap.wtd.Events)
		}
		ap.TocInfo["num_events"] = fmt.Sprintf("%d", l)
	}
}

func (ap *AppPageData) IsIdValid() bool {
	fileName := filepath.Join(ServerConfig.Paths.DataPath, ""+ap.ThingId+".json")
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		return false
	}
	return true
}

func (ap *AppPageData) Serialize() error {
	var fileName string
	var err error
	fileName = filepath.Join(ServerConfig.Paths.DataPath, ""+ap.ThingId+".json")
	if err = ap.wtd.Serialize(ap.ThingId, fileName); err != nil {
		Error.Printf("Serialize() for wtd id=%s, err=%s\n", ap.ThingId, err)
		return err
	}

	fileName = filepath.Join(ServerConfig.Paths.DataPath, ""+ap.ThingId+".meta.json")
	if err = ap.md.Serialize(ap.ThingId, fileName); err != nil {
		Error.Printf("Serialize() for md id=%s, err=%s\n", ap.ThingId, err)
		return err
	}

	return nil
}

func (ap *AppPageData) Deserialize() error {
	var fileName string
	var err error
	fileName = filepath.Join(ServerConfig.Paths.DataPath, ""+ap.ThingId+".json")
	ap.wtd = &WebThingDescription{}
	if err = ap.wtd.Deserialize(ap.ThingId, fileName); err != nil {
		Error.Printf("Deserialize() for wtd id=%s, err=%s\n", ap.ThingId, err)
		return err
	}

	fileName = filepath.Join(ServerConfig.Paths.DataPath, ""+ap.ThingId+".meta.json")
	ap.md = &GeneratorMetaData{}
	if err = ap.md.Deserialize(ap.ThingId, fileName); err != nil {
		Error.Printf("Deserialize() for md, id=%s, err=%s\n", ap.ThingId, err)
		return err
	}

	return nil
}

type AppErrorPageData struct {
	AppPageData
	Message string
}

func AppErrorServePage(w http.ResponseWriter, message string, id string) {
	templates, err := NewBasicHtmlTemplateSet("app_error.html.tpl")
	if err != nil {
		Error.Fatalf("Fatal error creating template set: %s\n", err)
	}

	data := &AppErrorPageData{
		AppPageData: AppPageData{
			PageData: PageData{
				Title: "Error",
			},
			ThingId: id,
		},
		Message: message,
	}

	if err = templates.ExecuteTemplate(w, "root", data); err != nil {
		Error.Printf("Error executing template: %s\n", err)
		w.WriteHeader(500)
		fmt.Fprint(w, "There was an internal error.")
	}
}

