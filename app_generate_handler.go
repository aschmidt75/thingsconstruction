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
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/fsouza/go-dockerclient"
	"github.com/gorilla/mux"
	"github.com/satori/go.uuid"
	"io/ioutil"
	"net/http"
	"os"
)

type appGenerateData struct {
	AppPageData
	Msg      string
	Accepted bool
}

func appGenerateNewPageData(id string) *appGenerateData {
	// read data from id
	data := &appGenerateData{
		AppPageData: AppPageData{
			PageData: PageData{
				Title: "Generate WoT code",
				InApp: true,
			},
			ThingId: id,
		},
		Accepted: false,
	}
	data.SetFeaturesFromConfig()
	if !data.IsIdValid() {
		return nil
	}
	if err := data.Deserialize(); err != nil {
		Error.Println(err)
		return nil
	}
	data.SetTocInfo()

	return data
}

func appGenerateNewResultPageData(id string) *appGenerateResultData {
	// read data from id
	data := &appGenerateResultData{
		AppPageData: AppPageData{
			PageData: PageData{
				Title: "Generate WoT code",
				InApp: true,
			},
			ThingId: id,
		},
	}
	data.SetFeaturesFromConfig()
	if !data.IsIdValid() {
		return nil
	}
	if err := data.Deserialize(); err != nil {
		Error.Println(err)
		return nil
	}
	data.SetTocInfo()

	return data
}

func AppGenerateHandleGet(w http.ResponseWriter, req *http.Request) {
	if ServerConfig.Features.App == false {
		http.Redirect(w, req, "/", 302)
		return
	}

	// check if id is valid
	vars := mux.Vars(req)
	data := appGenerateNewPageData(vars["id"])
	if data == nil {
		AppErrorServePage(w, "An error occurred while reading session data. Please try again.", vars["id"])
	}

	appGenerateServePage(w, data)

}

func AppGenerateDataHandleGet(w http.ResponseWriter, req *http.Request) {
	if ServerConfig.Features.App == false {
		w.WriteHeader(501)
		return
	}

	// check if id is valid
	vars := mux.Vars(req)
	id := vars["id"]

	data := appGenerateNewPageData(id)
	if data == nil {
		w.WriteHeader(500)
		fmt.Fprint(w, "Error deserializing session data")
		return
	}
	Debug.Printf("id=%s, wtd=%s\n", id, spew.Sdump(data.wtd))

	t, err := ReadGeneratorsConfig()
	if err != nil {
		Error.Printf("Unable to present generators. FIX CONFIG!\n")
		w.WriteHeader(500)
		fmt.Fprint(w, "Error loading generator data")
		return
	}

	pageData := struct {
		Wtd    *WebThingDescription `json:"wtd"`
		Target *AppGenTarget        `json:"target"`
	}{Wtd: data.wtd, Target: t.AppGenTargetById(data.md.SelectedGeneratorId)}

	b, err := json.Marshal(pageData)
	if err != nil {
		Error.Println(err)
		w.WriteHeader(500)
		fmt.Fprint(w, "Error marshaling data")
		return
	}

	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	w.Write(b)
}

func AppGenerateAcceptHandlePost(w http.ResponseWriter, req *http.Request) {
	if ServerConfig.Features.App == false {
		w.WriteHeader(http.StatusNotImplemented)
		return
	}

	err := req.ParseForm()
	if err != nil {
		Debug.Printf("Error parsing generate form: %s\n", err)
		appGenerateServePage(w, &appGenerateData{
			AppPageData: AppPageData{
				Message: "There was an error processing your data.",
			}})
	}
	formData := req.PostForm
	Debug.Printf(spew.Sdump(formData))

	// check if id is valid
	vars := mux.Vars(req)
	id := vars["id"]

	pageData := struct {
		Id    string
		Token string
	}{Id: id, Token: AppGenerateToken()}

	b, err := json.Marshal(pageData)
	if err != nil {
		Error.Println(err)
		w.WriteHeader(500)
		fmt.Fprint(w, "Error marshaling data")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	w.Write(b)
}

func AppGenerateHandlePost(w http.ResponseWriter, req *http.Request) {
	if ServerConfig.Features.App == false {
		w.WriteHeader(http.StatusNotImplemented)
		return
	}

	err := req.ParseForm()
	if err != nil {
		Debug.Printf("Error parsing generate form: %s\n", err)
		appGenerateServePage(w, &appGenerateData{
			AppPageData: AppPageData{
				Message: "There was an error processing your data.",
			}})
	}
	formData := req.PostForm
	Debug.Printf(spew.Sdump(formData))

	// check if id and token is valid
	vars := mux.Vars(req)
	id := vars["id"]
	_ = formData.Get("id")
	token := formData.Get("token")

	data := appGenerateNewPageData(id)

	if AppCheckToken(token) {
		if err := runModule(data); err != nil {
			Error.Println(err)
			var msg = fmt.Sprintf("An internal error occurred while generating your code. (%s)", err)
			data.AppPageData.Message = msg

			appGenerateServePage(w, &appGenerateData{
				AppPageData: AppPageData{
					Message: msg,
				}})

			return
		}

		http.Redirect(w, req, fmt.Sprintf("/app/%s/result", id), 302)
	} else {
		appGenerateServePage(w, &appGenerateData{
			AppPageData: AppPageData{
				Message: "There was an error processing your data.",
			}})

		return
	}
}

func runModule(data *appGenerateData) error {
	targets, err := ReadGeneratorsConfig()
	if err != nil || targets == nil {
		return errors.New("i1")
	}

	target := targets.AppGenTargetById(data.md.SelectedGeneratorId)
	Verbose.Printf("Using target=%#v", target)

	client, err := docker.NewClient("unix:///var/run/docker.sock")
	if err != nil {
		Error.Println(err)
		return errors.New("i2")
	}
	client.SkipServerVersionCheck = true

	Debug.Printf("cli=%#v\n", client)

	// create or renew i/o folders in basepath
	basePath, err := GetBasePathByThingId(data.ThingId)
	if err != nil {
		Error.Println(err)
		return errors.New("i3")
	}

	runId := uuid.Must(uuid.NewV4()).String()
	Debug.Printf("runId=%s", runId)

	inPath := fmt.Sprintf("%s/%s-in", basePath, runId)
	outPath := fmt.Sprintf("%s/%s-out", basePath, runId)
	err = os.Mkdir(inPath, 0777)
	if err != nil {
		Error.Println(err)
		return errors.New("i4")
	}
	err = os.Mkdir(outPath, 0777)
	if err != nil {
		Error.Println(err)
		return errors.New("i5")
	}

	hostMounts := make([]docker.HostMount, 2)
	hostMounts[0] = docker.HostMount{
		Target:   "/in",
		Source:   inPath,
		Type:     "bind",
		ReadOnly: false,
	}
	hostMounts[1] = docker.HostMount{
		Target:   "/out",
		Source:   outPath,
		Type:     "bind",
		ReadOnly: false,
	}

	opts := docker.CreateContainerOptions{
		Config: &docker.Config{
			Image:     target.ImageRepoTag,
			OpenStdin: true,
			StdinOnce: true,
		},
		HostConfig: &docker.HostConfig{
			Mounts: hostMounts,
		},
	}

	// Create the container, start the container
	container, err := client.CreateContainer(opts)
	if err != nil {
		Error.Println(err)
		return errors.New("i10")
	}

	Debug.Printf("container=%#v\n", container)

	err = client.StartContainer(container.ID, &docker.HostConfig{})
	if err != nil {
		Error.Println(err)
		return errors.New("i11")
	}

	// attach stdin, stdout and stderr to container.
	Debug.Printf("Started, attaching\n")

	// create a module request
	mr := NewModuleRequest(data.ThingId)

	// copy input files to temp stage.
	srcPath := fmt.Sprintf("%s/data.json", basePath)
	destPath := fmt.Sprintf("%s/%s-in/data.json", basePath, runId)

	tempData, err := ioutil.ReadFile(srcPath)
	if err != nil {
		Error.Println(err)
		return errors.New("i12")
	}
	// Write data to dst
	err = ioutil.WriteFile(destPath, tempData, 0600)
	if err != nil {
		Error.Println(err)
		return errors.New("i13")
	}
	// add input file to module request
	mr.AddInputFile("data.json")

	// for simulating stdin, we use a string reader with predefined content.
	var reader = mr.ShipRequest()
	var buf bytes.Buffer
	var buferr bytes.Buffer

	attachOpts := docker.AttachToContainerOptions{
		Container:    container.ID,
		Stdin:        true,
		Stdout:       true,
		Stderr:       true,
		InputStream:  reader,
		OutputStream: &buf,
		ErrorStream:  &buferr,
		Stream:       true,
		Logs:         true,
	}

	err = client.AttachToContainer(attachOpts)
	if err != nil {
		Error.Println(err)
		return errors.New("i12")
	}

	// Wait until container has finished. TODO: WaitContainerWithContext, timeout, ...
	exitCode, err := client.WaitContainer(container.ID)
	if err != nil {
		Error.Println(err)
		return errors.New("i13")
	}

	// dump some results.
	Debug.Printf("Exitcode=%#v\n", exitCode)
	if exitCode != 0 {
		Error.Printf("Module returned non-zero exit code: %d. Will not continue", exitCode)
		return errors.New("i14")
	}

	Debug.Println(buf.String())
	// save to file for later usage
	ioutil.WriteFile(fmt.Sprintf("%s/last-result.json", basePath), buf.Bytes(), 0640)
	ioutil.WriteFile(fmt.Sprintf("%s/last-result.id", basePath), []byte(runId), 0640)
	// Parse this reponse, construct data for web page.
	moduleResponse, err := ParseResponseFromModule(buf.Bytes())
	if err != nil {
		Error.Printf("Error in parsing response from module. CHECK MODULE. %s", err)
		return errors.New("i15")
	}
	Debug.Printf("%s", spew.Sdump(moduleResponse))
	if moduleResponse.Status == "error" {
		data.Message = fmt.Sprintf("Module reported an error while generating your code: %s // ID: %s", *moduleResponse.Message, data.ThingId)
		return errors.New("i16")
	}

	Debug.Printf("%s", spew.Sdump(moduleResponse.Files))
	// if we get have, things probably went well.
	//data.Files = moduleResponse.Files

	Verbose.Println(buferr.String())

	buf.Reset()
	buferr.Reset()

	return nil
}

func appGenerateServePage(w http.ResponseWriter, data *appGenerateData) {
	templates, err := NewBasicHtmlTemplateSet("app_generate.html.tpl", "app_generate_script.html.tpl")
	if err != nil {
		Error.Fatalf("Fatal error creating template set: %s\n", err)
	}

	err = templates.ExecuteTemplate(w, "root", data)
	if err != nil {
		Error.Printf("Error executing template: %s\n", err)
		w.WriteHeader(500)
		fmt.Fprint(w, "There was an internal error.")
	}

}
