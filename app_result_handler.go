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
	"archive/tar"
	"archive/zip"
	"bytes"
	"compress/gzip"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/shurcooL/github_flavored_markdown"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
)

type appGenerateResultData struct {
	AppPageData
	Msg              string
	Files            *ModuleResponseFiles
	URLPrefix        string
	CustomizationUrl string
}

func AppGenerateResultWtdHandleGet(w http.ResponseWriter, req *http.Request) {
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
	data.UpdateFeaturesFromContext(req.Context())

	b, err := json.MarshalIndent(data.wtd, "", "\t")
	if err != nil {
		Error.Println(err)
		w.WriteHeader(500)
		fmt.Fprint(w, "Error marshaling data")
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Content-Disposition", "attachment; filename=\""+data.ThingId+".json\"")
	w.WriteHeader(200)
	w.Write(b)
}

func AppGenerateResultAssetHandleGet(w http.ResponseWriter, req *http.Request) {
	if ServerConfig.Features.App == false {
		w.WriteHeader(501)
		return
	}

	// check if id is valid
	vars := mux.Vars(req)
	id := vars["id"]
	permalink := vars["permalink"]

	data := appGenerateNewResultPageData(id)
	data.UpdateFeaturesFromContext(req.Context())

	if data == nil {
		w.WriteHeader(500)
		fmt.Fprint(w, "Error deserializing session data")
		return
	}

	// load all results
	mr, err := loadResults(data)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, "Error deserializing generation result data")
		return
	}

	err = serveAssetFrom(data, mr, permalink, w, true)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, "Error serving generation result data")
		return
	}
}

func AppGenerateResultHandleGet(w http.ResponseWriter, req *http.Request) {
	if ServerConfig.Features.App == false {
		w.WriteHeader(501)
		return
	}

	// check if id is valid
	vars := mux.Vars(req)
	id := vars["id"]

	//
	data := &appGenerateResultData{
		AppPageData: AppPageData{
			PageData: PageData{
				Title: "Generate WoT code - Results",
				InApp: true,
			},
			ThingId: id,
		},
		URLPrefix: ServerConfig.Paths.URLPrefix,
	}
	data.SetFeaturesFromConfig()
	data.UpdateFeaturesFromContext(req.Context())

	if !data.IsIdValid() {
		Error.Printf("Generate: id not valid")
		appGenerateServePage(w, &appGenerateData{
			AppPageData: AppPageData{
				Message: "There was an error processing your data.",
			}})
		return
	}
	if err := data.Deserialize(); err != nil {
		Error.Println(err)
		appGenerateServePage(w, &appGenerateData{
			AppPageData: AppPageData{
				Message: "There was an error processing your data.",
			}})
		return
	}

	// load all results
	mr, err := loadResults(data)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, "Error deserializing generation result data")
		return
	}

	// do we have a customization for the current module?
	genId := data.md.SelectedGeneratorId
	appGenTargets, err := ReadGeneratorsConfig()
	if err == nil {
		appGenTarget := appGenTargets.AppGenTargetById(genId)
		if appGenTarget.CustomizationApp != "" {
			// find url
			customAppData, err := ServerConfig.GetCustomizationAppsDetailByName(appGenTarget.CustomizationApp)
			if err == nil {
				data.CustomizationUrl = fmt.Sprintf(customAppData.Entrypoint1, id)
			} else {
				Error.Printf("Could not get custom app url %s", err)
			}
		}
	} else {
		Error.Printf("Error reading generator config: %s", err)
	}

	data.Files = mr.Files
	appGenerateResultServePage(w, data)
}

func AppGenerateResultAssetViewHandleGet(w http.ResponseWriter, req *http.Request) {
	if ServerConfig.Features.App == false {
		w.WriteHeader(501)
		return
	}

	// check if id is valid
	vars := mux.Vars(req)
	id := vars["id"]
	permalink := vars["permalink"]

	Debug.Printf("Looking for id=%s, l=%s", id, permalink)
	data := appGenerateNewResultPageData(id)
	if data == nil {
		w.WriteHeader(500)
		fmt.Fprint(w, "Error deserializing session data")
		return
	}
	data.UpdateFeaturesFromContext(req.Context())

	// load all results
	mr, err := loadResults(data)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, "Error deserializing generation result data")
		return
	}

	err = serveAssetAsMDFrom(data, mr, permalink, w, false)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, "Error serving generation result data (view mode)")
		return
	}

}

func AppGenerateResultAssetArchiveHandleGet(w http.ResponseWriter, req *http.Request) {
	if ServerConfig.Features.App == false {
		w.WriteHeader(501)
		return
	}

	// check if id is valid
	vars := mux.Vars(req)
	id := vars["id"]
	format := vars["format"]

	Debug.Printf("Looking for id=%s, format=%s", id, format)
	data := appGenerateNewResultPageData(id)
	if data == nil {
		w.WriteHeader(500)
		fmt.Fprint(w, "Error deserializing session data")
		return
	}
	data.UpdateFeaturesFromContext(req.Context())

	// load all results
	mr, err := loadResults(data)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, "Error deserializing generation result data")
		return
	}

	if format == "zip" {
		err = serveAssetArchiveAsZIP(data, mr, w)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, "Error serving generation result data")
			return
		}
		return
	}
	if format == "tar" {
		err = serveAssetArchiveAsTAR(data, mr, w)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, "Error serving generation result data")
			return
		}
		return
	}
	if format == "targz" {
		err = serveAssetArchiveAsTARGZ(data, mr, w)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, "Error serving generation result data")
			return
		}
		return
	}

	w.WriteHeader(500)
	fmt.Fprint(w, "Unsupported format")
}

func loadResults(data *appGenerateResultData) (*ModuleResponse, error) {
	basePath, err := GetBasePathByThingId(data.ThingId)
	if err != nil {
		Error.Println(err)
		return nil, err
	}

	// look into the files "last-result.*" they must be present
	lastResultJsonFileName := fmt.Sprintf("%s/last-result.json", basePath)
	lastResultRunIdFileName := fmt.Sprintf("%s/last-result.id", basePath)

	_, err = os.Stat(lastResultJsonFileName)
	if err != nil {
		Error.Printf("Unable to find result files for id=%s", data.ThingId)
		return nil, err
	}
	_, err = os.Stat(lastResultRunIdFileName)
	if err != nil {
		Error.Printf("Unable to find result files (2) for id=%s", data.ThingId)
		return nil, err
	}

	b, err := ioutil.ReadFile(lastResultJsonFileName)
	res, err := ParseResponseFromModule(b)
	return res, nil
}

func readAssetFile(data *appGenerateResultData) ([]byte, error) {
	basePath, err := GetBasePathByThingId(data.ThingId)
	if err != nil {
		Error.Println(err)
		return nil, err
	}

	// look into the files "last-result.*" they must be present
	lastResultRunIdFileName := fmt.Sprintf("%s/last-result.id", basePath)

	_, err = os.Stat(lastResultRunIdFileName)
	if err != nil {
		Error.Printf("Unable to find result files (2) for id=%s", data.ThingId)
		return nil, err
	}
	b, err := ioutil.ReadFile(lastResultRunIdFileName)

	return b, err
}

func serveAssetFrom(data *appGenerateResultData, mr *ModuleResponse, permaLink string, w http.ResponseWriter, bAsFile bool) error {
	basePath, err := GetBasePathByThingId(data.ThingId)
	if err != nil {
		Error.Println(err)
		return err
	}

	b, err := readAssetFile(data)
	if err != nil {
		return err
	}

	// iterate over result files, locate permalink
	for _, file := range *mr.Files {
		if *file.Permalink == permaLink {

			// locate file within folder
			filePath := fmt.Sprintf("%s/%s-out/%s", basePath, string(b), file.FileName)
			b, err = ioutil.ReadFile(filePath)

			ct := "text/plain; charset=utf-8"

			if bAsFile {
				if file.ContentType != nil {
					ct = *file.ContentType
				}
				w.Header().Set("Content-Disposition", "attachment; filename=\""+file.FileName+"\"")
			} else {
				//				w.Header().Set("Content-Disposition", "inline")
			}
			w.Header().Set("Content-Type", ct)
			w.Header().Set("Content-Size", fmt.Sprintf("%d", len(b)))
			w.WriteHeader(200)
			w.Write(b)

			return nil
		}
	}

	return errors.New("unable to find file for permalink")
}

func serveAssetAsMDFrom(data *appGenerateResultData, mr *ModuleResponse, permaLink string, w http.ResponseWriter, bAsFile bool) error {
	basePath, err := GetBasePathByThingId(data.ThingId)
	if err != nil {
		Error.Println(err)
		return err
	}

	b, err := readAssetFile(data)
	if err != nil {
		return err
	}

	// iterate over result files, locate permalink
	for _, file := range *mr.Files {
		if *file.Permalink == permaLink {

			// locate file within folder
			filePath := fmt.Sprintf("%s/%s-out/%s", basePath, string(b), file.FileName)
			b, err = ioutil.ReadFile(filePath)

			templates, err := NewHtmlTemplateSet("root", "app_result_view.html.tpl")
			if err != nil {
				Error.Printf("Unable to create template set for viewing results: %s", err)
				return err
			}
			if *file.FileType == "Source Code" {
				bStr := string(b)
				wrappedStr := fmt.Sprintf("```%s\n%s\n```", *file.Language, bStr)
				b = []byte(wrappedStr)
			}

			context := &struct {
				MainContent template.HTML
			}{
				MainContent: template.HTML(github_flavored_markdown.Markdown(b)),
			}
			if err := templates.Execute(w, context); err != nil {
				Error.Printf("Unable to execute template for viewing results: %s", err)
				return err
			}
			return nil
		}
	}

	return errors.New("unable to find file for permalink or invalid content type")
}

func serveAssetArchiveAsZIP(data *appGenerateResultData, mr *ModuleResponse, w http.ResponseWriter) error {
	basePath, err := GetBasePathByThingId(data.ThingId)
	if err != nil {
		Error.Println(err)
		return err
	}

	// look into the files "last-result.*" they must be present
	lastResultRunIdFileName := fmt.Sprintf("%s/last-result.id", basePath)

	_, err = os.Stat(lastResultRunIdFileName)
	if err != nil {
		Error.Printf("Unable to find result files (2) for id=%s", data.ThingId)
		return err
	}
	bRunId, err := ioutil.ReadFile(lastResultRunIdFileName)
	Debug.Printf("runid=%s", bRunId)

	// pack as zip
	buf := new(bytes.Buffer)
	zw := zip.NewWriter(buf)
	defer zw.Close()

	// iterate over result files
	for _, file := range *mr.Files {
		f, err := zw.Create(file.FileName)
		if err != nil {
			Error.Printf("Error writing zip id=%s: %s", data.ThingId, err)
			return err
		}

		// write file into archive
		filePath := fmt.Sprintf("%s/%s-out/%s", basePath, string(bRunId), file.FileName)
		b, err := ioutil.ReadFile(filePath)
		if err != nil {
			Error.Printf("Error (2) writing zip id=%s: %s", data.ThingId, err)
			return err
		}
		_, err = f.Write(b)
		if err != nil {
			Error.Printf("Error (3) writing zip id=%s: %s", data.ThingId, err)
			return err
		}
	}
	err = zw.Close()
	if err != nil {
		Error.Printf("Error (4) writing zip id=%s: %s", data.ThingId, err)
		return err
	}

	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", "attachment; filename=\""+data.ThingId+".zip\"")
	w.WriteHeader(200)
	w.Write(buf.Bytes())

	return nil
}

func serveAssetArchiveAsTAR(data *appGenerateResultData, mr *ModuleResponse, w http.ResponseWriter) error {
	basePath, err := GetBasePathByThingId(data.ThingId)
	if err != nil {
		Error.Println(err)
		return err
	}

	// look into the files "last-result.*" they must be present
	lastResultRunIdFileName := fmt.Sprintf("%s/last-result.id", basePath)

	_, err = os.Stat(lastResultRunIdFileName)
	if err != nil {
		Error.Printf("Unable to find result files (2) for id=%s", data.ThingId)
		return err
	}
	bRunId, err := ioutil.ReadFile(lastResultRunIdFileName)
	Debug.Printf("runid=%s", bRunId)

	//
	buf := new(bytes.Buffer)

	tw := tar.NewWriter(buf)
	defer tw.Close()

	serveAssetArchiveAsTARRaw(data, mr, w, tw, basePath, bRunId)

	w.Header().Set("Content-Type", "application/tar")
	w.Header().Set("Content-Disposition", "attachment; filename=\""+data.ThingId+".tar\"")
	w.WriteHeader(200)
	w.Write(buf.Bytes())

	return nil
}

func serveAssetArchiveAsTARGZ(data *appGenerateResultData, mr *ModuleResponse, w http.ResponseWriter) error {
	basePath, err := GetBasePathByThingId(data.ThingId)
	if err != nil {
		Error.Println(err)
		return err
	}

	// look into the files "last-result.*" they must be present
	lastResultRunIdFileName := fmt.Sprintf("%s/last-result.id", basePath)

	_, err = os.Stat(lastResultRunIdFileName)
	if err != nil {
		Error.Printf("Unable to find result files (2) for id=%s", data.ThingId)
		return err
	}
	bRunId, err := ioutil.ReadFile(lastResultRunIdFileName)
	Debug.Printf("runid=%s", bRunId)

	//
	buf := new(bytes.Buffer)

	gw := gzip.NewWriter(buf)
	defer gw.Close()

	tw := tar.NewWriter(buf)
	defer tw.Close()

	if err = serveAssetArchiveAsTARRaw(data, mr, w, tw, basePath, bRunId); err != nil {
		Error.Printf("Unable to server as tar: %s", err)
		return err
	}

	w.Header().Set("Content-Type", "application/gzip")
	w.Header().Set("Content-Disposition", "attachment; filename=\""+data.ThingId+".tar.gz\"")
	w.WriteHeader(200)
	w.Write(buf.Bytes())

	return nil
}

func serveAssetArchiveAsTARRaw(data *appGenerateResultData, mr *ModuleResponse, w http.ResponseWriter, tw *tar.Writer, basePath string, bRunId []byte) error {
	// iterate over result files
	for _, file := range *mr.Files {
		filePath := fmt.Sprintf("%s/%s-out/%s", basePath, string(bRunId), file.FileName)
		b, err := ioutil.ReadFile(filePath)
		if err != nil {
			Error.Printf("Error (1) writing tar id=%s: %s", data.ThingId, err)
			return err
		}

		hdr := &tar.Header{
			Name: file.FileName,
			Mode: 0640,
			Size: int64(len(b)),
		}
		if err := tw.WriteHeader(hdr); err != nil {
			Error.Printf("Error writing (2) tar id=%s: %s", data.ThingId, err)
			return err
		}
		if _, err := tw.Write(b); err != nil {
			Error.Printf("Error writing (3) tar id=%s: %s", data.ThingId, err)
			return err
		}
	}
	err := tw.Close()
	if err != nil {
		Error.Printf("Error (4) writing tar id=%s: %s", data.ThingId, err)
		return err
	}
	return nil
}

func appGenerateResultServePage(w http.ResponseWriter, data *appGenerateResultData) {
	templates, err := NewBasicHtmlTemplateSet("app_result.html.tpl", "app_result_script.html.tpl")
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
