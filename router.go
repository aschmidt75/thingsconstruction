//  ThingsConstruction, a code generator for WoT-based models
//  Copyright (C) 2017,2018  @aschmidt75
//
//  This program is free software: you can redistribute it and/or modify
//  it under the terms of the GNU Affero General Public License as published
//  by the Free Software Foundation, either version 3 of the License, or
//  (at your option) any later version.
//
//  This program is distributed in the hope that it will be useful,
//  but WITHOUT ANY WARRANTY; without even the implied warranty of
//  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//  GNU Affero General Public License for more details.
//
//  You should have received a copy of the GNU Affero General Public License
//  along with this program.  If not, see <http://www.gnu.org/licenses/>.
//
//  This program is dual-licensed. For commercial licensing options, please
//  contact the author(s).
//

//
package main

import (
	"context"
	"fmt"
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
	Route{"DisableCookiesTracking", "GET", "/donottrack", doNotTrackHandler},
	Route{"ResetCookies", "GET", "/cookieconsentreset", cookieConsentResetHandler},
	Route{"Index", "GET", "/index.html", IndexHandler},
	Route{"Index", "GET", "/", IndexHandler},
	Route{"StaticPage", "GET", "/{page}.html", StaticPageHandler},
	Route{"Blog", "GET", "/blog", BlogIndexHandler},
	Route{"Blog", "GET", "/blog/data", BlogIndexJSONHandler},
	Route{"BlogPage", "GET", "/blog/{page}", MarkdownBlogHandler},
	Route{"ModuleInfo", "GET", "/module/{id}", ModulePageHandler},
	Route{"ModuleInfo", "GET", "/modules/data", ModuleDataHandler},
	Route{"AppCreateThing", "GET", "/app", AppCreateThingHandleGet},
	Route{"AppCreateThing", "POST", "/app", AppCreateThingHandlePost},
	Route{"AppCreateThing", "POST", "/app/fromtemplate", AppCreateThingFromTemplateHandlePost},
	Route{"AppCreateThing", "GET", "/app/{id}", AppCreateThingHandleGet},
	Route{"AppCreateThing", "POST", "/app/{id}", AppCreateThingHandlePost},
	Route{"AppCreateThing", "DELETE", "/app/{id}", AppCreateThingHandleDelete},
	Route{"AppChooseFramework", "GET", "/app/{id}/framework", AppChooseFrameworkHandleGet},
	Route{"AppChooseFramework", "POST", "/app/{id}/framework", AppChooseFrameworkHandlePost},
	Route{"AppManageProperties", "GET", "/app/{id}/properties", AppManagePropertiesHandleGet},
	Route{"AppManageProperties", "GET", "/app/{id}/properties/data", AppManagePropertiesDataHandleGet},
	Route{"AppManageProperties", "POST", "/app/{id}/properties", AppManagePropertiesHandlePost},
	Route{"AppManageActions", "GET", "/app/{id}/actions", AppManageActionsHandleGet},
	Route{"AppManageActions", "GET", "/app/{id}/actions/data", AppManageActionsDataHandleGet},
	Route{"AppManageActions", "POST", "/app/{id}/actions", AppManageActionsHandlePost},
	Route{"AppManageEvents", "GET", "/app/{id}/events", AppManageEventsHandleGet},
	Route{"AppManageEvents", "GET", "/app/{id}/events/data", AppManageEventsDataHandleGet},
	Route{"AppManageEvents", "POST", "/app/{id}/events", AppManageEventsHandlePost},
	Route{"AppGenerate", "GET", "/app/{id}/generate", AppGenerateHandleGet},
	Route{"AppGenerate", "GET", "/app/{id}/generate/data", AppGenerateDataHandleGet},
	Route{"AppGenerate", "POST", "/app/{id}/generate/accept", AppGenerateAcceptHandlePost},
	Route{"AppGenerate", "POST", "/app/{id}/generate", AppGenerateHandlePost},
	Route{"AppResult", "GET", "/app/{id}/result", AppGenerateResultHandleGet},
	Route{"AppResult", "GET", "/app/{id}/result/wtd", AppGenerateResultWtdHandleGet},
	Route{"AppResult", "GET", "/app/{id}/result/asset/{permalink}", AppGenerateResultAssetHandleGet},
	Route{"AppResult", "GET", "/app/{id}/result/assetview/{permalink}", AppGenerateResultAssetViewHandleGet},
	Route{"AppResult", "GET", "/app/{id}/result/asset-archive/{format}", AppGenerateResultAssetArchiveHandleGet},
	Route{"Feedback", "GET", "/feedback", FeedbackHandleGet},
	Route{"Feedback", "POST", "/feedback", FeedbackHandlePost},
	Route{"FeedbackVote", "POST", "/feedback/vote", FeedbackVoteHandlePost},
	Route{"FeedbackQuick", "POST", "/feedback/q", FeedbackQuickHandlePost},
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)

	// add asset paths
	staticPaths := []string{"js", "css", "fonts", "img", "webfonts"}
	for _, staticPath := range staticPaths {
		p := fmt.Sprintf("/%s/", staticPath)
		d := fmt.Sprintf("%s/%s", ServerConfig.Paths.AssetPath, staticPath)
		router.PathPrefix(p).Handler(http.StripPrefix(p, http.FileServer(http.Dir(d))))
	}

	router.PathPrefix("/dimg/").Handler(http.StripPrefix("/dimg/", http.FileServer(http.Dir(
		fmt.Sprintf("%s/dimg", ServerConfig.Paths.MDPagesPath)))))

	// add application routes
	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = addNoCacheHeaders(handler)
		handler = filterTooBigPayloads(handler)
		handler = cookieProcessingHandler(handler)
		handler = logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	router.NotFoundHandler = http.HandlerFunc(notFoundHandler)

	return router
}

func notFoundHandler(w http.ResponseWriter, req *http.Request) {
	templates, err := NewBasicHtmlTemplateSet("_404.html.tpl")
	if err != nil {
		Error.Println(err)
		w.WriteHeader(404)
		fmt.Fprint(w, "The page you're looking for has not been found.")
	}

	//
	data := &appGenerateResultData{
		AppPageData: AppPageData{
			PageData: PageData{
				Title: "Not found",
				InApp: true,
			},
		},
	}
	data.SetFeaturesFromConfig()
	data.UpdateFeaturesFromContext(req.Context())

	err = templates.ExecuteTemplate(w, "root", data)
	if err != nil {
		Error.Println(err)
		w.WriteHeader(404)
		fmt.Fprint(w, "The page you're looking for has not been found.")
	}
}

func doNotTrackHandler(w http.ResponseWriter, req *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "cookieconsent_status",
		Value: "deny",
	})
	http.SetCookie(w, &http.Cookie{
		Name:  "_ga",
		Value: "",
	})
	http.SetCookie(w, &http.Cookie{
		Name:  "_gid",
		Value: "",
	})
	http.Redirect(w, req, "/", http.StatusMovedPermanently)
}

func cookieConsentResetHandler(w http.ResponseWriter, req *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "cookieconsent_status",
		Value: "",
	})
	http.Redirect(w, req, "/", http.StatusMovedPermanently)
}

func logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		Debug.Printf(
			"> %s %s %s",
			r.Method,
			r.RequestURI,
			name,
		)
		inner.ServeHTTP(w, r)

		Verbose.Printf(
			"< %s %s %s %s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}

// Insert no-cache elements into http header
func addNoCacheHeaders(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=UTF-8")
		w.Header().Set("Cache-control", "no-cache")
		w.Header().Set("Cache-control", "no-store")
		w.Header().Set("Pragma", "no-cache")
		w.Header().Set("Expires", "0")
		if ServerConfig.Http.ReferrerPolicy != "" {
			w.Header().Set("Referrer-Policy", ServerConfig.Http.ReferrerPolicy)
		}

		if ServerConfig.Http.ContentSecurityPolicy != "" {
			w.Header().Set("Content-Security-Policy", ServerConfig.Http.ContentSecurityPolicy)
		}
		w.Header().Set("X-Frame-Options", "DENY")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("X-XSS-Protection", "1; mode=block")

		// forward
		inner.ServeHTTP(w, r)
	})
}

func cookieProcessingHandler(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// check cookie consent
		var bConsentOk = false
		cookieConsentStatus, err := r.Cookie("cookieconsent_status")

		// with OptIn, we either get "allow" or "dismiss"

		if err != nil {
			// unable to find cookie
			bConsentOk = false
		} else {
			if cookieConsentStatus != nil && cookieConsentStatus.Value == "allow" {
				bConsentOk = true
			}
		}

		//Debug.Printf("consent ok=%b", bConsentOk)

		if !bConsentOk {
			// we may not set cookies, b/c user did not answer cookie consent yet or answered with "deny".
			// make following code in chain turn off features that would set cookies
			Debug.Printf("Turning off cookies")
			ctx := context.WithValue(r.Context(), "tc-nocookies", true)
			inner.ServeHTTP(w, r.WithContext(ctx))
		} else {
			// we may set cookies.

			// check for feature param. if given, make it permanent as a cookie
			featureActivation := r.URL.Query().Get("with-feature")
			if featureActivation != "" {
				// set cookie
				http.SetCookie(w, &http.Cookie{
					Name:  "tc-feature",
					Value: featureActivation,
				})
			}

			c, _ := r.Cookie("tc-feature")
			if c != nil {
				ctx := context.WithValue(r.Context(), "tc-feature", c.Value)
				inner.ServeHTTP(w, r.WithContext(ctx))
			} else {
				inner.ServeHTTP(w, r)
			}

		}
	})
}
func filterTooBigPayloads(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.Body = http.MaxBytesReader(w, r.Body, 4096)
		inner.ServeHTTP(w, r)
	})
}

// fav icon of http://www.iconspedia.com/icon/things-digital-icon-22104.html
func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/img/favicon.ico")
}
