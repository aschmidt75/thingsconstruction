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
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
	"strings"
	"time"
)

type BlogPageMetaData struct {
	Title            string
	Author           string
	Date             string
	DateTime         time.Time
	Tags             []string
	Abstract         string
	originalFileName string
}

type BlogPage struct {
	Name     string
	MetaData *BlogPageMetaData
	Content  []byte
}

type BlogMetaData struct {
	NumPages    int
	Authors     map[string][]*BlogPage
	PagesByDate []BlogPage
	AllTags     map[string][]*BlogPage
}

type BlogPages struct {
	MetaData *BlogMetaData
	Pages    map[string]*BlogPage
}

func (bp *BlogPages) addBlogPage(p *BlogPage) error {
	bp.Pages[p.Name] = p

	// update all-over metadata
	bp.MetaData.NumPages = bp.MetaData.NumPages + 1

	key := p.MetaData.Author
	if len(key) > 0 {
		arr, ok := bp.MetaData.Authors[key]
		if ok == false {
			arr = make([]*BlogPage, 0)
			bp.MetaData.Authors[key] = arr
		}
		arr = append(arr, p)
	}

	for _, tag := range p.MetaData.Tags {
		arr, ok := bp.MetaData.AllTags[tag]
		if ok == false {
			arr = make([]*BlogPage, 0)
			bp.MetaData.AllTags[tag] = arr
		}
		arr = append(arr, p)
	}

	//Debug.Printf("%+v\n", bp.MetaData.Authors)
	//Debug.Printf("%+v\n", bp.MetaData.AllTags)
	return nil
}

// Reloads all pages from origins
func (bp *BlogPages) Reload() error {
	Debug.Printf("Reloading blog pages...")
	var err error
	for _, page := range bp.Pages {
		if e := page.Reload(); e != nil {
			err = e
		}
	}
	if err != nil {
		Error.Printf("Error reloading (some) blog pages: %s\n", err)
	}
	return err
}

func NewBlogPages() *BlogPages {
	b := &BlogPages{
		MetaData: &BlogMetaData{},
		Pages:    make(map[string]*BlogPage),
	}

	// link pages by authors and tags
	b.MetaData.Authors = make(map[string][]*BlogPage)
	b.MetaData.AllTags = make(map[string][]*BlogPage)

	return b
}

// Given path as directory, reads all files from it,
// stores it under BlogPages struct type.
func NewBlogPagesFromPath(path string) (*BlogPages, error) {
	// treat path as dir, read all from dir
	fileInfos, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	b := NewBlogPages()

	for _, fileInfo := range fileInfos {
		Debug.Printf("Found %s\n", fileInfo.Name())
		pathName := filepath.Join(path, fileInfo.Name())
		bp, err := NewBlogPageFromPath(pathName)
		if err != nil {
			Verbose.Printf("Cannot load markdown page %s: %s\n", pathName, err)
			continue
		}

		b.addBlogPage(bp)
	}

	return b, nil
}

func NewBlogPageFromPath(path string) (*BlogPage, error) {
	bp := &BlogPage{}
	bp.Name = filepath.Base(path)
	bp.MetaData = &BlogPageMetaData{}
	bp.MetaData.originalFileName = path

	if err := bp.Reload(); err != nil {
		return nil, err
	}
	return bp, nil
}

func (bp *BlogPage) Reload() error {
	path := bp.MetaData.originalFileName
	input, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	// split data into frontmatter and content
	lines := strings.Split(string(input), "\n")
	frontMatter := ""
	content := ""
	bFrontMatter := false
	bFrontMatterComplete := false
	for _, line := range lines {
		if !bFrontMatterComplete && strings.Index(line, "---") == 0 {
			if bFrontMatter == false {
				bFrontMatter = true
			} else {
				bFrontMatterComplete = true
				bFrontMatter = false
			}
		} else {
			if bFrontMatter {
				frontMatter = frontMatter + line + "\n"
			} else {
				content = content + line + "\n"
			}

		}
	}

	if len(frontMatter) > 0 {
		err = yaml.Unmarshal([]byte(frontMatter), bp.MetaData)
		if err != nil {
			return err
		}
		if len(bp.MetaData.Date) > 0 {
			bp.MetaData.DateTime, err = time.Parse("02.01.2006", bp.MetaData.Date)
			if err != nil {
				Verbose.Printf("Error parsing date in %s\n", path)
			}
		}
	}

	bp.Content = []byte(content)

	return nil
}
