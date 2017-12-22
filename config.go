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
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	Http struct {
		Port int `yaml:"port"`
	}
	Logging struct {
		Debug   string `yaml:"debug"`
		Verbose string `yaml:"verbose"`
		Error   string `yaml:"error"`
	}
	Paths struct {
		AssetPath          string `yaml:"assets"`
		StaticPagesPath    string `yaml:"staticpages"`
		MDPagesPath        string `yaml:"mdpages"`
		FeedbackPath       string `yaml:"feedback"`
		GeneratorsDataPath string `yaml:"generators"`
		DataPath           string `yaml:"data"`
	}
	Features struct {
		Blog    bool `yaml:"blog"`
		App     bool `yaml:"app"`
		Contact bool `yaml:"contact"`
	}
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
