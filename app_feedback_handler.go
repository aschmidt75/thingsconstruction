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
package main

import (
	"fmt"
	"net/http"
	"path/filepath"
	"io/ioutil"
	"os"
	"time"
)

func makeTimestampStr() string {
	return fmt.Sprintf("feedback-%d.txt", (time.Now().UnixNano() % 1e6 / 1e3))
}
func AppFeedbackQuickHandlePost(w http.ResponseWriter, req *http.Request) {
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
	fbPath := filepath.Join(ServerConfig.Paths.FeedbackPath,makeTimestampStr())
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
