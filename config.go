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

type Config struct {
	Http struct {
		Port int `yaml:"port"`
		ContentSecurityPolicy string `yaml:"csp"`
	}
	Logging struct {
		Debug   string `yaml:"debug"`
		Verbose string `yaml:"verbose"`
		Error   string `yaml:"error"`
	}
	Paths struct {
		AssetPath          string `yaml:"assets"`
		StaticPagesPath    string `yaml:"staticpages"`
		ModulePagesPath    string `yaml:"modulepages"`
		MDPagesPath        string `yaml:"mdpages"`
		FeedbackPath       string `yaml:"feedback"`
		GeneratorsDataPath string `yaml:"generators"`
		DataPath           string `yaml:"data"`
		URLPrefix          string `yaml:"urlprefix"`
	}
	Features struct {
		Blog    bool `yaml:"blog"`
		App     bool `yaml:"app"`
		Contact bool `yaml:"contact"`
		Twitter bool `yaml:"twitter"`
		LinkedIn bool `yaml:"linkedin"`
		GitHub  bool `yaml:"github"`
		Analytics  bool `yaml:"analytics"`
		Shariff  bool `yaml:"shariff"`
		VoteForGenerators  bool `yaml:"vote_generators"`
		Flattr  bool `yaml:"flattr"`
	}
	StaticTexts struct {
		LinkedInUrl string `yaml:"linkedin_url"`
		GitHubUrl string `yaml:"github_url"`
		TwitterUrl string `yaml:"twitter_url"`
		CopyrightLine string `yaml:"copyrightline"`
		Notices string `yaml:"notices"`
		FlattrId  string `yaml:"flattrid"`
		FlattrUser  string `yaml:"flattruser"`
	}
	Docker struct {
		UserConfig string `yaml:"userConfig"`
	}
	VoteGenerators map[string]string  `yaml:"vote_generators"`
}

// Read yaml configuration from file (fromPath),
// unmarshal into Config struct, return it.
func NewConfig(fromPath string) (*Config, error) {
	b, err := ioutil.ReadFile(fromPath)
	if err != nil {
		Error.Printf("Error reading config file, %s: %s\n", fromPath, err)
		return nil, err
	}

	res := &Config{}
	err = yaml.Unmarshal(b, res)

	return res, err
}
