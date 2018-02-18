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
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strings"
)

type ModuleResponseFile struct {
	Permalink   *string `json:"permalink"`
	FileName    string  `json:"filename"`
	Description *string `json:"desc"`
	ContentType *string `json:"ct"`
	FileType    *string `json:"type"`
	Language    *string `json:"language"`
}
type ModuleResponseFiles []ModuleResponseFile

type ModuleResponse struct {
	Status  string               `json:"status"`
	Message *string              `json:"msg"`
	Files   *ModuleResponseFiles `json:"files"`
}

type ModuleRequestFile struct {
	FileName    string `json:"filename"`
	ContentType string `json:"ct"`
	FileType    string `json:"type"`
}
type ModuleRequestFiles []ModuleRequestFile

type ModuleRequest struct {
	ThingId string              `json:"thingid"`
	Files   *ModuleRequestFiles `json:"files"`
}

func NewModuleRequest(id string) *ModuleRequest {
	res := &ModuleRequest{
		ThingId: id,
		Files:   &ModuleRequestFiles{},
	}
	return res
}

func (mr *ModuleRequest) AddInputFile(filePath string) {
	*mr.Files = append(*mr.Files, ModuleRequestFile{
		FileName:    filePath,
		ContentType: "application/json",
		FileType:    "thingdescription",
	})
}

func (mr *ModuleRequest) ShipRequest() *strings.Reader {
	b, err := json.Marshal(mr)
	if err != nil {
		return strings.NewReader("")
	}

	return strings.NewReader(string(b))
}

func MakePermaLink(mrf *ModuleResponseFile) string {
	tmpStr := fmt.Sprintf("tc-%s", mrf.FileName)

	s := sha256.Sum256([]byte(tmpStr))
	sb := []byte(s[:])
	return fmt.Sprintf("%s", hex.EncodeToString(sb))
}

func ParseResponseFromModule(b []byte) (*ModuleResponse, error) {
	res := &ModuleResponse{}
	err := json.Unmarshal(b, res)
	if err == nil && res.Files != nil {
		// create permalinks over files
		for idx := 0; idx < len(*res.Files); idx++ {
			file := &(*res.Files)[idx]
			file.Permalink = new(string)
			*file.Permalink = MakePermaLink(file)
		}

	}
	return res, err
}

type modulePageData struct {
	PageData
	HtmlOutput template.HTML
}

func ModulePageHandler(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	pageName := vars["page"]

	// look up module by name
	bp := "sample-module.html"
	ok := true
	if ok {
		bp = filepath.Join(ServerConfig.Paths.ModulePagesPath, bp)
		Debug.Printf("serving module page %s\n", bp)

		tplBytes, err := ioutil.ReadFile(bp)
		if err != nil {
			Error.Printf("Error reading page by name %s\n", pageName)
			ServeNotFound(w, req)
			return
		}

		modulePagesServePage(w, modulePageData{
			PageData: PageData{
				Title: pageName, // TODO
			},
			HtmlOutput: template.HTML(tplBytes),
		})
	} else {
		ServeNotFound(w, req)
	}

}

var ModulePagesTemplates *template.Template

func initializeModuleTemplates() {
	if ModulePagesTemplates == nil {
		Debug.Printf("Initializing templates for module pages")
		var err error
		ModulePagesTemplates, err = NewBasicHtmlTemplateSet("staticpage.html.tpl", "staticpage_script.html.tpl")
		if err != nil {
			Error.Fatalf("Fatal error creating template set: %s\n", err)
		}
	}
}

func modulePagesServePage(w http.ResponseWriter, data modulePageData) {
	initializeModuleTemplates()
	data.SetFeaturesFromConfig()

	err := ModulePagesTemplates.ExecuteTemplate(w, "root", data)
	if err != nil {
		Error.Printf("Error executing template: %s\n", err)
		w.WriteHeader(500)
		fmt.Fprint(w, "There was an internal error.")
	}

}
