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
//
//    This program is dual-licensed. For commercial licensing options, please
//    contact the author(s).
//
package main

import (
	"html/template"
	"io/ioutil"
	"path/filepath"
)

// Adds a a file to the template set
func AddTemplateByFileName(template *template.Template, fileName string) error {
	tplBytes, err := ioutil.ReadFile(filepath.Join("./templates/", fileName))
	if err != nil {
		return err
	}

	_, err = template.New(fileName).Parse(string(tplBytes))
	if err != nil {
		return err
	}

	return nil
}

// adds one or more files as templates
func AddTemplateByFileNames(template *template.Template, fileNames ...string) error {
	for _, fileName := range fileNames {
		err := AddTemplateByFileName(template, fileName)
		if err != nil {
			return err
		}
	}
	return nil
}

// adds default templates head, header, footer to the tempalte set
func AddDefaultTemplates(template *template.Template) error {
	return AddTemplateByFileNames(template, "_head.html.tpl", "_header.html.tpl", "_footer.html.tpl")
}

// creates a new template, pre-filled with name default templates
// to build the web page
func NewHtmlTemplateSet(name string, fileNames ...string) (*template.Template, error) {
	templates := template.New(name)
	fileName := name + ".html.tpl"
	err := AddTemplateByFileName(templates, fileName)
	if err != nil {
		return templates, err
	}
	for _, fileName := range fileNames {
		err = AddTemplateByFileName(templates, fileName)
		if err != nil {
			return templates, err
		}
	}
	err = AddDefaultTemplates(templates)
	return templates, err
}

// Includes root template, names template set as root.
// adds empty_script template if script template not present.
func NewBasicHtmlTemplateSet(fileNames ...string) (*template.Template, error) {
	templates := template.New("root")
	fileName := "root.html.tpl"
	err := AddTemplateByFileName(templates, fileName)
	if err != nil {
		return templates, err
	}
	for _, fileName := range fileNames {
		err = AddTemplateByFileName(templates, fileName)
		if err != nil {
			return templates, err
		}
	}
	err = AddDefaultTemplates(templates)
	if err != nil {
		return templates, err
	}
	// add empty script template if not present
	if templates.Lookup("script") == nil {
		err = AddTemplateByFileNames(templates, "_empty_script.html.tpl")
	}
	return templates, err
}
