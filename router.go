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
	"time"
)


type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{"FavIcon", "GET", "/favicon", faviconHandler},
	Route{"Index", "GET", "/index.html", IndexHandler},
	Route{"Index", "GET", "/", IndexHandler},
	Route{"About", "GET", "/{page}.html", StaticPageHandler},
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)

	staticPaths := []string{"js", "css", "fonts", "img"}
	for _, staticPath := range staticPaths {
		p := "/" + staticPath + "/"
		d := "./static/" + staticPath
		router.PathPrefix(p).Handler(http.StripPrefix(p, http.FileServer(http.Dir(d))))
	}

	// add application routes
	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = webNoCache(handler)
		handler = logger(handler, route.Name) // Log request

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		Debug.Printf(
			"> %s\t%s\t%s",
			r.Method,
			r.RequestURI,
			name,
		)
		inner.ServeHTTP(w, r)

		Verbose.Printf(
			"< %s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}


// Insert no-cache elements into http header
func webNoCache(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=UTF-8")
		w.Header().Set("Cache-control", "no-cache")
		w.Header().Set("Cache-control", "no-store")
		w.Header().Set("Pragma", "no-cache")
		w.Header().Set("Expires", "0")

		// forward
		inner.ServeHTTP(w, r)
	})
}

// fav icon of http://www.iconspedia.com/icon/things-digital-icon-22104.html
func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/img/favicon.ico")
}