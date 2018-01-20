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
	Name        string   `json:name`
	Type        string   `json:type`
	MaxLength   *int     `json:maxlength,omitempty`
	Min         *float64 `json:maxlength,omitempty`
	Max         *float64 `json:maxlength,omitempty`
	Description *string  `json:description,omitempty`
}
type WebThingProperties []WebThingProperty

type WebThingAction struct {
	Name        string  `json:name`
	Description *string `json:description,omitempty`
}
type WebThingActions []WebThingAction
type WebThingEvent struct {
	Name        string  `json:name`
	Description *string `json:description,omitempty`
}
type WebThingEvents []WebThingEvent

type WebThingDescription struct {
	Name        string              `json:name`
	Type        string              `json:type`
	Description *string             `json:description,omitempty`
	Properties  *WebThingProperties `json:properties,omitempty`
	Actions     *WebThingActions    `json:actions,omitempty`
	Events      *WebThingEvents     `json:events,omitempty`
}

func (wtd *WebThingDescription) NewProperties() {
	wtd.Properties = &WebThingProperties{}
	*wtd.Properties = make([]WebThingProperty, 0, 3)
}

func (wtd *WebThingDescription) AppendProperty(p WebThingProperty) {
	*wtd.Properties = append(*wtd.Properties, p)
}

func (wtd *WebThingDescription) NewActions() {
	wtd.Actions = &WebThingActions{}
	*wtd.Actions = make([]WebThingAction, 0, 3)
}

func (wtd *WebThingDescription) AppendAction(a WebThingAction) {
	*wtd.Actions = append(*wtd.Actions, a)
}

func (wtd *WebThingDescription) NewEvents() {
	wtd.Events = &WebThingEvents{}
	*wtd.Events = make([]WebThingEvent, 0, 3)
}

func (wtd *WebThingDescription) AppendEvent(a WebThingEvent) {
	*wtd.Events = append(*wtd.Events, a)
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
