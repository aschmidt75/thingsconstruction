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
	"encoding/json"
	"io/ioutil"
)

type WebThingDescription struct {
	Name        string `json:name`
	Type        string `json:type`
	Description string `json:description`
}

func (self *WebThingDescription) Serialize(id string, fileName string) error {
	b, err := json.Marshal(self)
	if err != nil {
		return err
	}

	Debug.Printf("%s %s", fileName, b)
	return ioutil.WriteFile(fileName, b, 0640)
}

func (self *WebThingDescription) Deserialize(id string, fileName string) error {
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		Debug.Printf("Error reading %s: %s", fileName, err)
		return err
	}

	err = json.Unmarshal(b, self)
	if err != nil {
		Debug.Printf("Error parsing %s: %s", fileName, err)
		Debug.Println(string(b))
	}
	return err
}
