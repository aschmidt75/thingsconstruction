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
	"encoding/json"
	"github.com/dustin/go-humanize"
	"github.com/gorilla/mux"
	"github.com/shurcooL/github_flavored_markdown"
	"html/template"
	"net/http"
	"regexp"
	"sort"
	"strings"
	"time"
)

type tagChipData struct {
	TagName string
	Active  bool
}

type blogPostChronoData struct {
	Title         string
	Date          time.Time
	DateFormatted string
	DateElapsed   string
	Name          string
	Tags          []string
	Abstract      string
	HasVideo      bool
}

type blogContentData struct {
	PageData

	MetaData       *BlogPageMetaData
	BlogMetaData   *BlogMetaData
	TagChipData    []tagChipData
	AllPostsChrono []blogPostChronoData
	HtmlOutput     template.HTML
	VimeoID        string
	YoutubeID      string
}

type blogOverviewData struct {
	PageData
	BlogMetaData   *BlogMetaData
	TagChipData    []tagChipData
	AllPostsChrono []blogPostChronoData
}

func BlogIndexHandler(w http.ResponseWriter, req *http.Request) {
	if ServerConfig.Features.Blog == false {
		http.Redirect(w, req, "/", 302)
		return
	}

	Blog.Reload()

	data := blogOverviewData{
		PageData: PageData{
			Title:  "Blog Index",
			InBlog: true,
			Robots: true,
		},
		BlogMetaData:   Blog.MetaData,
		TagChipData:    collectTagChipData(Blog.MetaData, nil),
		AllPostsChrono: collectAllPostsChrono(Blog),
	}
	data.SetFeaturesFromConfig()
	data.UpdateFeaturesFromContext(req.Context())

	blogServeOverviewPage(w, data)
}

func BlogIndexJSONHandler(w http.ResponseWriter, req *http.Request) {
	if ServerConfig.Features.Blog == false {
		w.WriteHeader(501)
		return
	}

	Blog.Reload()

	b, err := json.Marshal(collectAllPostsChrono(Blog))
	if err != nil {
		w.WriteHeader(500)
		return

	}
	w.WriteHeader(200)
	w.Write(b)
}

func MarkdownBlogHandler(w http.ResponseWriter, req *http.Request) {
	if ServerConfig.Features.Blog == false {
		http.Redirect(w, req, "/", 302)
		return
	}

	vars := mux.Vars(req)
	pageName := vars["page"]

	Blog.Reload()

	// look up page by name
	extMatch, _ := regexp.MatchString(".*\\.md$", pageName)
	if !extMatch {
		pageName = pageName + ".md"
	}
	bp, ok := Blog.Pages[pageName]
	if ok {
		markDown := github_flavored_markdown.Markdown(bp.Content)
		r := strings.NewReplacer("<ul>", "<ul class=\"browser-default\">")
		htmlStr := template.HTML(r.Replace(string(markDown)))

		data := blogContentData{
			PageData: PageData{
				Title:           bp.MetaData.Title,
				MetaDescription: bp.MetaData.Abstract,
				InBlog:          true,
				Robots:          true,
			},
			MetaData:       bp.MetaData,
			BlogMetaData:   Blog.MetaData,
			TagChipData:    collectTagChipData(Blog.MetaData, bp.MetaData),
			AllPostsChrono: collectAllPostsChrono(Blog),
			HtmlOutput:     htmlStr,
			VimeoID:        bp.MetaData.VimeoID,
			YoutubeID:      bp.MetaData.YoutubeID,
		}
		data.SetFeaturesFromConfig()
		data.UpdateFeaturesFromContext(req.Context())

		Debug.Printf("%#v", data)
		blogServePage(w, data)

		UseCaseCounter.Increment("blogpost-viewed", pageName)

	} else {
		Verbose.Printf("Unable to find page by name %s\n", pageName)
		w.WriteHeader(404)
	}

}

func collectTagChipData(blogMetaData *BlogMetaData, blogPageMetaData *BlogPageMetaData) []tagChipData {
	// collect and sort all tags
	t := make([]tagChipData, 0)
	for tagName := range blogMetaData.AllTags {
		bActive := false

		if blogPageMetaData != nil {
			// look up this page for active tags
			for _, tagName0 := range blogPageMetaData.Tags {
				if tagName0 == tagName {
					bActive = true
				}
			}
		}
		t = append(t, tagChipData{
			TagName: tagName,
			Active:  bActive,
		})
	}
	sort.Slice(t, func(i, j int) bool {
		return t[i].TagName < t[j].TagName
	})

	return t
}

func collectAllPostsChrono(blog *BlogPages) []blogPostChronoData {
	// collect and sort all (recent) posts
	cr := make([]blogPostChronoData, 0)
	for name, post := range blog.Pages {
		cr = append(cr, blogPostChronoData{
			Title:         post.MetaData.Title,
			Name:          name,
			Date:          post.MetaData.DateTime,
			DateFormatted: post.MetaData.DateTime.Format("Mon, Jan 2 2006"),
			DateElapsed:   humanize.Time(post.MetaData.DateTime),
			Tags:          post.MetaData.Tags,
			Abstract:      post.MetaData.Abstract,
			HasVideo:      (post.MetaData.YoutubeID != "") || (post.MetaData.VimeoID != ""),
		})
	}
	sort.Slice(cr, func(i, j int) bool {
		return cr[j].Date.Before(cr[i].Date)
	})

	return cr
}

func blogServePage(w http.ResponseWriter, data blogContentData) {

	templates, err := NewHtmlTemplateSet("root", "blog.html.tpl", "blog_script.html.tpl")
	if err != nil {
		Error.Fatalf("Fatal error creating template set: %s\n", err)
	}

	err = templates.ExecuteTemplate(w, "root", data)
	if err != nil {
		Verbose.Printf("Error executing template: %s\n", err)
	}

}

func blogServeOverviewPage(w http.ResponseWriter, data blogOverviewData) {

	templates, err := NewHtmlTemplateSet("root", "blog_overview.html.tpl", "blog_script.html.tpl")
	if err != nil {
		Error.Fatalf("Fatal error creating template set: %s\n", err)
	}

	err = templates.ExecuteTemplate(w, "root", data)
	if err != nil {
		Verbose.Printf("Error executing template: %s\n", err)
	}

}
