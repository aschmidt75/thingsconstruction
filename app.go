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
	"encoding/json"
	"errors"
	"fmt"
	"github.com/satori/go.uuid"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"sync"
	"time"
)

type AppTokens struct {
	tokens map[string]time.Time
	mux    sync.Mutex
}

var CurrentAppTokens AppTokens

type AppError struct {
	Message string
}

func (ae *AppError) Error() string {
	return ae.Message
}

type GeneratorMetaData struct {
	SelectedGeneratorId string `json:"genid"`
}

func (gmd *GeneratorMetaData) Serialize(id string, fileName string) error {
	b, err := json.Marshal(gmd)
	if err != nil {
		return err
	}

	Debug.Printf("%s %s", fileName, b)
	return ioutil.WriteFile(fileName, b, 0640)
}

func (gmd *GeneratorMetaData) Deserialize(id string, fileName string) error {
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		Debug.Printf("Error reading %s: %s", fileName, err)
		return err
	}

	err = json.Unmarshal(b, gmd)
	if err != nil {
		Debug.Printf("Error parsing %s: %s", fileName, err)
		Debug.Println(string(b))
	}
	return err
}

func AppTokensNew() {
	CurrentAppTokens.mux.Lock()
	defer CurrentAppTokens.mux.Unlock()
	CurrentAppTokens.tokens = make(map[string]time.Time)
}

// create a new token and store it
func AppGenerateToken() string {
	CurrentAppTokens.mux.Lock()
	defer CurrentAppTokens.mux.Unlock()

	t := uuid.Must(uuid.NewV4()).String()

	CurrentAppTokens.tokens[t] = time.Now().Add(time.Hour * 1)

	return t
}

func AppCheckToken(t string) bool {

	exp, ok := CurrentAppTokens.tokens[t]
	if !ok {
		return false
	}
	if exp.Before(time.Now()) {
		return false
	}

	// todo:sweep through all tokens, remove expired ones.
	CurrentAppTokens.mux.Lock()
	defer CurrentAppTokens.mux.Unlock()

	return true
}

type AppPageData struct {
	PageData
	ThingId string
	TocInfo map[string]string
	wtd     *WebThingDescription
	md      *GeneratorMetaData
	Message string
}

func (ap *AppPageData) SetTocInfo() {
	ap.TocInfo = map[string]string{}

	if ap.wtd != nil {
		var l int
		l = 0
		if ap.wtd.Properties != nil {
			l = len(ap.wtd.Properties)
		}
		ap.TocInfo["num_properties"] = fmt.Sprintf("%d", l)

		l = 0
		if ap.wtd.Actions != nil {
			l = len(ap.wtd.Actions)
		}
		ap.TocInfo["num_actions"] = fmt.Sprintf("%d", l)

		l = 0
		if ap.wtd.Events != nil {
			l = len(ap.wtd.Events)
		}
		ap.TocInfo["num_events"] = fmt.Sprintf("%d", l)
	}
}

// checks if ap.ThingId has valid format and is stored
// on local data path
func (ap *AppPageData) IsIdValid() bool {
	id := "" + ap.ThingId
	if len(id) == 0 {
		return false
	}
	r := regexp.MustCompile("[0-9a-zA-Z_-]+")
	if r.MatchString(id) == false {
		Debug.Printf("Invalid ID form: %s", id)
		return false
	}

	dirName := filepath.Join(ServerConfig.Paths.DataPath, ""+ap.ThingId)
	if s, err := os.Stat(dirName); os.IsNotExist(err) || s.IsDir() == false {
		return false
	}

	fileName := filepath.Join(dirName, "data.json")
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		return false
	}
	return true
}

func GetBasePathByThingId(thingId string) (string, error) {
	var err error

	dirName := filepath.Join(ServerConfig.Paths.DataPath, ""+thingId)
	s, err := os.Stat(dirName)
	if os.IsNotExist(err) {
		return "", errors.New("unable to access thing data")
	} else {
		if s.IsDir() == false {
			return "", errors.New("unable to access thing data (file exist but is not a dir)")
		}
	}
	return dirName, nil
}

func (ap *AppPageData) Serialize() error {
	var err error

	dirName := filepath.Join(ServerConfig.Paths.DataPath, ""+ap.ThingId)
	s, err := os.Stat(dirName)
	if os.IsNotExist(err) {
		// make dir
		err = os.Mkdir(dirName, 0777)
		if err != nil {
			return err
		}
	} else {
		if s.IsDir() == false {
			return errors.New("unable to store thing data (file exist but is not a dir)")
		}
	}

	var fileName string
	fileName = filepath.Join(dirName, "data.json")
	if err = ap.wtd.Serialize(ap.ThingId, fileName); err != nil {
		Error.Printf("Serialize() for wtd id=%s, err=%s\n", ap.ThingId, err)
		return err
	}

	fileName = filepath.Join(dirName, "meta.json")
	if err = ap.md.Serialize(ap.ThingId, fileName); err != nil {
		Error.Printf("Serialize() for md id=%s, err=%s\n", ap.ThingId, err)
		return err
	}

	return nil
}

func (ap *AppPageData) Deserialize() error {
	if !ap.IsIdValid() {
		Error.Printf("Deserialize() for wtd id=%s: not a valid id\n", ap.ThingId)
		return errors.New("not a valid id")
	}

	var fileName string
	var err error
	fileName = filepath.Join(ServerConfig.Paths.DataPath, ""+ap.ThingId, "data.json")
	ap.wtd = &WebThingDescription{}
	if err = ap.wtd.Deserialize(ap.ThingId, fileName); err != nil {
		Error.Printf("Deserialize() for wtd id=%s, err=%s\n", ap.ThingId, err)
		return err
	}

	fileName = filepath.Join(ServerConfig.Paths.DataPath, ""+ap.ThingId, "meta.json")
	ap.md = &GeneratorMetaData{}
	if err = ap.md.Deserialize(ap.ThingId, fileName); err != nil {
		Error.Printf("Deserialize() for md, id=%s, err=%s\n", ap.ThingId, err)
		return err
	}

	return nil
}

func (ap *AppPageData) Delete() error {
	if !ap.IsIdValid() {
		Error.Printf("Delete() for wtd id=%s: not a valid id\n", ap.ThingId)
		return errors.New("not a valid id")
	}

	dir := filepath.Join(ServerConfig.Paths.DataPath, ""+ap.ThingId)
	Verbose.Printf("Deleteing dir upon request: %s", dir)

	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		Verbose.Printf("Deleteing: %s", name)
		err = os.RemoveAll(filepath.Join(dir, name))
		if err != nil {
			return err
		}
	}

	err = os.RemoveAll(dir)
	if err != nil {
		return err
	}

	return nil
}

type AppErrorPageData struct {
	AppPageData
	Message string
}

func AppErrorServePage(w http.ResponseWriter, message string, id string) {
	templates, err := NewBasicHtmlTemplateSet("app_error.html.tpl")
	if err != nil {
		Error.Fatalf("Fatal error creating template set: %s\n", err)
	}

	data := &AppErrorPageData{
		AppPageData: AppPageData{
			PageData: PageData{
				Title: "Error",
			},
			ThingId: id,
		},
		Message: message,
	}

	if err = templates.ExecuteTemplate(w, "root", data); err != nil {
		Error.Printf("Error executing template: %s\n", err)
		w.WriteHeader(500)
		fmt.Fprint(w, "There was an internal error.")
	}
}
