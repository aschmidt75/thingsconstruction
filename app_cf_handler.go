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
}

type AppGenTargetArray []AppGenTarget

type AppGenTargets struct {
	Targets AppGenTargetArray `yaml:"targets"`
}

type appGenParamsData struct {
	PageData
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

func AppGenParamsHandleGet(w http.ResponseWriter, req *http.Request) {
	if ServerConfig.Features.App == false {
		http.Redirect(w, req, "/", 302)
		return
	}

	vars := mux.Vars(req)
	Verbose.Printf("Vars: %#v\n", vars)

	t, err := readGeneratorsConfig()

	if err != nil {
		t = &AppGenTargets{}
	}
	var data = &appGenParamsData{
		PageData: PageData{
			Title: "THNGS:CONSTR - Choose Embedded Development Framework",
		},
		NumGenerators: 1,
		AppGenTargets: t,
	}
	data.SetFeaturesFromConfig()
	data.InApp = true

	appGenParamsServePage(w, *data)
}

func appGenParamsServePage(w http.ResponseWriter, data appGenParamsData) {
	templates, err := NewBasicHtmlTemplateSet("app_cf.html.tpl", "app_cf_script.html.tpl")
	if err != nil {
		Error.Printf("Fatal error creating template set: %s\n", err)
	}

	if err = templates.ExecuteTemplate(w, "root", data); err != nil {
		Verbose.Printf("Error executing template: %s\n", err)
	}

}
