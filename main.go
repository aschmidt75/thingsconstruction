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
	"log"
	"net/http"
	"fmt"
)

var Debug *log.Logger
var Verbose *log.Logger
var Error *log.Logger
var ServerConfig *Config

func main() {
	Debug, Verbose, Error = InitializeBasicLogging()

	var err error
	ServerConfig, err = NewConfig("./etc/config.yaml")
	if err != nil {
		panic(err)
	} else {
		Debug, Verbose, Error = InitializeLogging(ServerConfig)
		Debug.Printf("%#v\n", ServerConfig)
	}

	router := NewRouter()
	Debug.Printf("router=%#v\n", router)

	srv := &http.Server{
		Addr:    	fmt.Sprintf(":%d", ServerConfig.Http.Port),
		Handler: 	router,
	}
	Debug.Printf("srv=%#v\n", srv)

	Verbose.Printf("Starting Server")
	err = srv.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
