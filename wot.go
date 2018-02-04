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

type WebThingProperty struct {
	Name        string   `json:"-"`
	Type        string   `json:"type"`
	MaxLength   *int     `json:"maxlength,omitempty"`
	Min         *float64 `json:"min,omitempty"`
	Max         *float64 `json:"max,omitempty"`
	Description *string  `json:"description,omitempty"`
}

type WebThingAction struct {
	Name        string  `json:"-"`
	Description *string `json:"description,omitempty"`
}

type WebThingEvent struct {
	Name        string  `json:"-"`
	Description *string `json:"description,omitempty"`
}

type WebThingDescription struct {
	Name        string                      `json:"name"`
	Type        string                      `json:"type"`
	Description *string                     `json:"description,omitempty"`
	Properties  map[string]WebThingProperty `json:"properties,omitempty"`
	Actions     map[string]WebThingAction   `json:"actions,omitempty"`
	Events      map[string]WebThingEvent    `json:"events,omitempty"`
}

func (wtd *WebThingDescription) NewProperties() {
	wtd.Properties = make(map[string]WebThingProperty)
}

func (wtd *WebThingDescription) AppendProperty(p WebThingProperty) {
	wtd.Properties[p.Name] = p
}

func (wtd *WebThingDescription) NewActions() {
	wtd.Actions = make(map[string]WebThingAction)
}

func (wtd *WebThingDescription) AppendAction(a WebThingAction) {
	wtd.Actions[a.Name] = a
}

func (wtd *WebThingDescription) NewEvents() {
	wtd.Events = make(map[string]WebThingEvent)
}

func (wtd *WebThingDescription) AppendEvent(a WebThingEvent) {
	wtd.Events[a.Name] = a
}

func (wtd *WebThingDescription) Serialize(id string, fileName string) error {
	b, err := json.Marshal(wtd)
	if err != nil {
		return err
	}

	Debug.Printf("%s %s", fileName, b)
	return ioutil.WriteFile(fileName, b, 0640)
}

func (wtd *WebThingDescription) Deserialize(id string, fileName string) error {
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		Debug.Printf("Error reading %s: %s", fileName, err)
		return err
	}

	err = json.Unmarshal(b, wtd)
	if err != nil {
		Debug.Printf("Error parsing %s: %s", fileName, err)
		Debug.Println(string(b))
	}
	return err
}
