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
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func FeedbackQuickHandlePost(w http.ResponseWriter, req *http.Request) {
	req.Body = http.MaxBytesReader(w, req.Body, 1024)
	err := req.ParseForm()
	if err != nil {
		Debug.Printf("Error parsing feedback form: %s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "text/plain")

		fmt.Fprint(w, "There was an error at the server.")
		return
	}
	cff := req.PostForm

	what := cff.Get("cff_feedback_what")
	Verbose.Printf("%s\n", what)

	// Save feedback somewhere
	fbPath := filepath.Join(ServerConfig.Paths.FeedbackPath, makeTimestampStr())
	if err := ioutil.WriteFile(fbPath, []byte(what), os.FileMode(400)); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "text/plain")

		fmt.Fprint(w, "I'm sorry, an error occured saving your feedback. Please try again later.")
		return
	}

	// send back a single line
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")

	fmt.Fprint(w, "Thank you for your feedback.")
}

func FeedbackHandlePost(w http.ResponseWriter, req *http.Request) {
	req.Body = http.MaxBytesReader(w, req.Body, 1024)

	if ServerConfig.Features.Contact == false {
		http.Redirect(w, req, "/", 302)
		return
	}

	err := req.ParseForm()
	if err != nil {
		Debug.Printf("Error parsing feedback form: %s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "text/plain")

		fmt.Fprint(w, "There was an error at the server.")
		return
	}
	cff := req.PostForm
	//Debug.Printf("feedback: %s", spew.Sdump(cff))

	data := struct {
		Category  string `json:"cat"`
		Content   string `json:"content"`
		Timestamp string `json:"ts"`
	}{
		Category:  cff.Get("fbf_what"),
		Content:   cff.Get("fbf_text"),
		Timestamp: fmt.Sprintf("%s", time.Now()),
	}
	Debug.Printf("feedback: %s", spew.Sdump(data))

	b, err := json.Marshal(data)
	Debug.Printf("%s\n", b)

	// Save feedback somewhere
	fbPath := filepath.Join(ServerConfig.Paths.FeedbackPath, makeTimestampStr2())
	if err := ioutil.WriteFile(fbPath, b, os.FileMode(400)); err != nil {
		Error.Printf("%s", err)

		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "text/plain")

		fmt.Fprint(w, "I'm sorry, an error occured saving your feedback. Please try again later.")
		return
	}
	Verbose.Printf("Feedback: written to %s\n", fbPath)

	// send back a single line
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")

	fmt.Fprint(w, "Thank you for your feedback.")
}

func FeedbackHandleGet(w http.ResponseWriter, req *http.Request) {
	if ServerConfig.Features.Contact == false {
		http.Redirect(w, req, "/", 302)
		return
	}

	data := &appEntryData{
		AppPageData: AppPageData{
			PageData: PageData{
				Title:     "Your feedback",
				InContact: true,
			},
		},
	}
	data.SetFeaturesFromConfig()

	templates, err := NewBasicHtmlTemplateSet("feedback.html.tpl", "feedback_script.html.tpl")
	if err != nil {
		Error.Fatalf("Fatal error creating template set: %s\n", err)
	}

	err = templates.ExecuteTemplate(w, "root", data)
	if err != nil {
		Error.Printf("Error executing template: %s\n", err)
	}
}

func FeedbackVoteHandlePost(w http.ResponseWriter, req *http.Request) {
	req.Body = http.MaxBytesReader(w, req.Body, 1024)

	if ServerConfig.Features.VoteForGenerators == false {
		http.Redirect(w, req, "/", 302)
		return
	}

	err := req.ParseForm()
	if err != nil {
		Debug.Printf("Error parsing feedback form: %s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	voteform := req.PostForm

	fbPath := filepath.Join(ServerConfig.Paths.FeedbackPath, makeTimestampStr3())
	file, err := os.OpenFile(fbPath, os.O_CREATE|os.O_RDWR, os.FileMode(400))
	if err != nil {
		Error.Printf("%s", err)

		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "text/plain")

		fmt.Fprint(w, "I'm sorry, an error occured saving your feedback. Please try again later.")
		return
	}
	defer file.Close()
	for key, name := range ServerConfig.VoteGenerators {
		value := voteform.Get(key)
		if value != "" {
			fmt.Fprintf(file, "%s;%s;%s\r\n", key, name, value)
		}
	}

	Verbose.Printf("Vote: written to %s\n", fbPath)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")

	fmt.Fprint(w, "Thank you for your vote.")

}

func makeTimestampStr() string {
	return fmt.Sprintf("feedback-%d.txt", (time.Now().UnixNano()))
}

func makeTimestampStr2() string {
	return fmt.Sprintf("feedback-%d.json", (time.Now().UnixNano()))
}

func makeTimestampStr3() string {
	return fmt.Sprintf("vote-%d.csv", (time.Now().UnixNano()))
}
