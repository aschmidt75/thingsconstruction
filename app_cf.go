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
	"github.com/gorilla/mux"
	"net/http"
)

type appGenTarget struct {
	Id        string
	ShortDesc string
	Desc      string
	Tags      []string
}

type appGenParamsData struct {
	PageData
	NumGenerators int
	AppGenTargets []appGenTarget
}

func AppGenParamsHandleGet(w http.ResponseWriter, req *http.Request) {
	if ServerConfig.Features.App == false {
		http.Redirect(w,req,"/", 302)
		return
	}


	vars := mux.Vars(req)
	Verbose.Printf("Vars: %#v\n", vars)

	// TODO: This from config file
	var t = make([]appGenTarget, 3)
	t[0] = appGenTarget{
		Id:        "1",
		ShortDesc: "Creates a HTTP Rest API using JSON, for the Arduino Frameworks and Ethernet Adapters",
		Desc:      "",
		Tags:      []string{"framework:arduino", "proto:http", "app:json", "conn:ethernet"},
	}
	t[1] = appGenTarget{
		Id:        "2",
		ShortDesc: "Creates a HTTP Rest API using JSON, for the ESP8266 Arduino Frameworks",
		Desc:      "",
		Tags:      []string{"framework:arduino", "proto:http", "app:json", "conn:wifi", "mcu:esp8266"},
	}
	t[2] = appGenTarget{
		Id:        "3",
		ShortDesc: "Creates a HTTP Rest API using MessagePack, for ARM MBed OS 5 compatible stuff, with WiFi",
		Desc:      "",
		Tags:      []string{"framework:ARM Mbed", "proto:http", "app:messagepack", "conn:wifi"},
	}

	var data = &appGenParamsData{
		PageData: PageData{
			Title: "THNGS:CONSTR - Choose Embedded Development Framework",
		},
		NumGenerators: 1,
		AppGenTargets: t,
	}
	data.SetFeaturesFromConfig()
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
