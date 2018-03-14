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
	"gopkg.in/yaml.v2"
	"io/ioutil"
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
	ImageRepoTag string                `yaml:"repotag"`
	Image        string                `yaml:"image"`
	ShortDesc    string                `yaml:"shortdesc"`
	Desc         string                `yaml:"desc"`
	Tags         AppStringArray        `yaml:"tags"`
	Dependencies AppGenDependencyArray `yaml:"dependencies"`
	CodeGenInfo  string                `yaml:"codegeninfo"`
}

type AppGenTargetArray []AppGenTarget

type AppGenTargets struct {
	Targets AppGenTargetArray `yaml:"targets"`
}

func (agt *AppGenTargets) AppGenTargetById(id string) *AppGenTarget {
	for _, target := range agt.Targets {
		if id == target.Id {
			return &target
		}
	}
	return nil
}

func ReadGeneratorsConfig() (*AppGenTargets, error) {
	b, err := ioutil.ReadFile(ServerConfig.Paths.GeneratorsDataPath)
	if err != nil {
		return nil, err
	}

	var t = &AppGenTargets{}
	if err := yaml.Unmarshal(b, t); err != nil {
		Error.Printf("Error reading generators config file. Check YAML: %s", err)
		return t, err
	}

	return t, nil
}
