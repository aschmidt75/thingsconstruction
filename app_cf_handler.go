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
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/http"
)

type AppStringArray []string

type AppGenDependency struct {
	Name      string `yaml:"name"`
	URL       string `yaml:"url"`
	Copyright string `yaml:"copyright"`
	License   string `yaml:"license"`
	Info      string `yaml:"info"`
}
type AppGenDependencyArray []AppGenDependency

type AppGenTarget struct {
	Id           string                `yaml:"id"`
	Image        string                `yaml:"image"`
	ShortDesc    string                `yaml:"shortdesc"`
	Desc         string                `yaml:"desc"`
	Tags         AppStringArray        `yaml:"tags"`
	Dependencies AppGenDependencyArray `yaml:"dependencies"`
	CodeGenInfo  string				   `yaml:"codegeninfo"`
}

type AppGenTargetArray []AppGenTarget

type AppGenTargets struct {
	Targets AppGenTargetArray `yaml:"targets"`
}

type appGenParamsData struct {
	AppPageData
	NumGenerators int
	AppGenTargets *AppGenTargets
}

func readGeneratorsConfig() (*AppGenTargets, error) {
	b, err := ioutil.ReadFile(ServerConfig.Paths.GeneratorsDataPath)
	if err != nil {
		return nil, err
	}

	var t = &AppGenTargets{}
	if err := yaml.Unmarshal(b, t); err != nil {
		Error.Printf("Error reading generators config file. Check YAML: %s", err)
		return t, err
	}

	Debug.Printf("targets=%s\n", spew.Sdump(t))
	return t, nil
}

func AppChooseFrameworkHandleGet(w http.ResponseWriter, req *http.Request) {
	if ServerConfig.Features.App == false {
		http.Redirect(w, req, "/", 302)
		return
	}

	vars := mux.Vars(req)
	id := vars["id"]

	// Todo: Check id, must be valid

	t, err := readGeneratorsConfig()

	if err != nil {
		t = &AppGenTargets{}
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

	appGenParamsServePage(w, *data)
}

func AppChooseFrameworkHandlePost(w http.ResponseWriter, req *http.Request) {
	if ServerConfig.Features.App == false {
		http.Redirect(w, req, "/", 302)
		return
	}

	err := req.ParseForm()
	if err != nil {
		Debug.Printf("Error parsing create thing form: %s\n", err)
		appCreateThingServePage(w, appEntryData{Msg: "There was an error processing your data."})
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
	if data.Deserialize() != nil {
		AppErrorServePage(w, "An error occurred while reading session data. Please try again.", id)
		return
	}
	Debug.Printf("id=%s, wtd=%s\n", spew.Sdump(data.wtd))

	// write generator to meta data file
	data.md = &GeneratorMetaData{
		SelectedGeneratorId: cfs,
	}
	if data.Serialize() != nil {
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
