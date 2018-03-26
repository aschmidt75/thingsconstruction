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
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/fsouza/go-dockerclient"
	"github.com/gorilla/mux"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type ModuleResponseFile struct {
	Permalink   *string `json:"permalink"`
	FileName    string  `json:"filename"`
	Description *string `json:"desc"`
	ContentType *string `json:"ct"`
	FileType    *string `json:"type"`
	Language    *string `json:"language"`
}
type ModuleResponseFiles []ModuleResponseFile

type ModuleResponse struct {
	Status  string               `json:"status"`
	Message *string              `json:"msg"`
	Files   *ModuleResponseFiles `json:"files"`
}

type ModuleRequestFile struct {
	FileName    string `json:"filename"`
	ContentType string `json:"ct"`
	FileType    string `json:"type"`
}
type ModuleRequestFiles []ModuleRequestFile

type ModuleRequest struct {
	ThingId string              `json:"thingid"`
	Files   *ModuleRequestFiles `json:"files"`
}

type ModuleMetaData struct {
	Status  string `json:"status"`
	Message string `json:"msg"`
	Details *interface{}

	Files ModuleResponseFiles
}

func NewModuleRequest(id string) *ModuleRequest {
	res := &ModuleRequest{
		ThingId: id,
		Files:   &ModuleRequestFiles{},
	}
	return res
}

func (mr *ModuleRequest) AddInputFile(filePath string) {
	*mr.Files = append(*mr.Files, ModuleRequestFile{
		FileName:    filePath,
		ContentType: "application/json",
		FileType:    "thingdescription",
	})
}

func (mr *ModuleRequest) ShipRequest() *strings.Reader {
	b, err := json.Marshal(mr)
	if err != nil {
		return strings.NewReader("")
	}

	return strings.NewReader(string(b))
}

func MakePermaLink(mrf *ModuleResponseFile) string {
	tmpStr := fmt.Sprintf("tc-%s", mrf.FileName)

	s := sha256.Sum256([]byte(tmpStr))
	sb := []byte(s[:])
	return fmt.Sprintf("%s", hex.EncodeToString(sb))
}

func ParseResponseFromModule(b []byte) (*ModuleResponse, error) {
	res := &ModuleResponse{}
	err := json.Unmarshal(b, res)
	if err == nil && res.Files != nil {
		// create permalinks over files
		for idx := 0; idx < len(*res.Files); idx++ {
			file := &(*res.Files)[idx]
			file.Permalink = new(string)
			*file.Permalink = MakePermaLink(file)
		}

	}
	return res, err
}

type modulePageData struct {
	PageData
	HtmlOutput template.HTML
}

func ModulePageHandler(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	moduleId := vars["id"]

	// start container to get module spec page
	content, err := getModuleSpecContent(moduleId)
	if err != nil {
		Error.Printf("Error reading module content for id=%s, err=%s\n", moduleId, err)
		ServeNotFound(w, req)
		return
	}

	modulePagesServePage(w, modulePageData{
		PageData: PageData{
			Title: "Module specification",
		},
		HtmlOutput: template.HTML(content),
	})
}

func ModuleDataHandler(w http.ResponseWriter, req *http.Request) {
	targets, err := ReadGeneratorsConfig()
	if err != nil || targets == nil {
		Error.Println(err)
		w.WriteHeader(500)
		fmt.Fprint(w, "Error reading modules data")
		return
	}

	b, err := json.Marshal(targets)
	if err != nil {
		Error.Println(err)
		w.WriteHeader(500)
		fmt.Fprint(w, "Error marshaling modules data")
		return
	}

	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	w.Write(b)
}

func getModuleSpecContent(moduleId string) ([]byte, error) {
	targets, err := ReadGeneratorsConfig()
	if err != nil || targets == nil {
		return nil, err
	}

	target := targets.AppGenTargetById(moduleId)
	if target == nil {
		return nil, errors.New("no module for id")
	}
	Verbose.Printf("Using target=%#v", target)

	client, err := docker.NewClient("unix:///var/run/docker.sock")
	if err != nil {
		return nil, err
	}
	client.SkipServerVersionCheck = true

	Debug.Printf("cli=%#v\n", client)

	outPath := fmt.Sprintf("%s/%s-out", ServerConfig.Paths.ModulePagesPath, moduleId)
	if _, err := os.Stat(outPath); err != nil {
		if os.IsNotExist(err) {
			if err = os.Mkdir(outPath, 0777); err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	outPath, err = filepath.Abs(outPath)
	if err != nil {
		return nil, err
	}

	hostMounts := make([]docker.HostMount, 1)
	hostMounts[0] = docker.HostMount{
		Target:   "/out",
		Source:   outPath,
		Type:     "bind",
		ReadOnly: false,
	}

	opts := docker.CreateContainerOptions{
		Config: &docker.Config{
			Image:     target.ImageRepoTag,
			OpenStdin: false,
			StdinOnce: false,
		},
		HostConfig: &docker.HostConfig{
			Mounts:  hostMounts,
			CapDrop: []string{"all"},
			CapAdd:  []string{"setuid", "setgid"},
		},
	}

	if ServerConfig.Docker.UserConfig != "" {
		opts.Config.User = ServerConfig.Docker.UserConfig
	}

	// Create the container, start the container
	container, err := client.CreateContainer(opts)
	if err != nil {
		return nil, err
	}

	Debug.Printf("container=%#v\n", container)

	if err = client.StartContainer(container.ID, &docker.HostConfig{}); err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	var buferr bytes.Buffer

	attachOpts := docker.AttachToContainerOptions{
		Container:    container.ID,
		Stdin:        false,
		Stdout:       true,
		Stderr:       true,
		OutputStream: &buf,
		ErrorStream:  &buferr,
		Stream:       true,
		Logs:         true,
	}

	if err = client.AttachToContainer(attachOpts); err != nil {
		return nil, err
	}

	// Wait until container has finished. TODO: WaitContainerWithContext, timeout, ...
	exitCode, err := client.WaitContainer(container.ID)
	if err != nil {
		return nil, err
	}

	// dump some results.
	Debug.Printf("Exitcode=%#v\n", exitCode)
	if exitCode != 0 {
		Error.Printf("Module returned non-zero exit code: %d. Will not continue", exitCode)
		return nil, errors.New("Non-zero exit code.")
	}

	var md ModuleMetaData
	if err := json.Unmarshal(buf.Bytes(), &md); err != nil {
		return nil, err
	}

	Debug.Println(spew.Sdump(md))

	// go through files, seek a module-spec type..
	for _, file := range md.Files {
		if *file.FileType == "module-spec" && *file.ContentType == "text/html" {
			return ioutil.ReadFile(filepath.Join(outPath, file.FileName))
		}
	}

	//	Debug.Println(buf.String())
	//	Debug.Println(buferr.String())

	return nil, errors.New("no suitable module spec found")
}

var ModulePagesTemplates *template.Template

func initializeModuleTemplates() {
	if ModulePagesTemplates == nil {
		Debug.Printf("Initializing templates for module pages")
		var err error
		ModulePagesTemplates, err = NewBasicHtmlTemplateSet("staticpage.html.tpl", "staticpage_script.html.tpl")
		if err != nil {
			Error.Fatalf("Fatal error creating template set: %s\n", err)
		}
	}
}

func modulePagesServePage(w http.ResponseWriter, data modulePageData) {
	initializeModuleTemplates()
	data.SetFeaturesFromConfig()

	err := ModulePagesTemplates.ExecuteTemplate(w, "root", data)
	if err != nil {
		Error.Printf("Error executing template: %s\n", err)
		w.WriteHeader(500)
		fmt.Fprint(w, "There was an internal error.")
	}

}
