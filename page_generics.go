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

// Generic Page Data, valid for all pages
// template data embeds this struct
type PageData struct {
	Title string

	Feature       map[string]bool
	CopyrightLine string

	InBlog    bool
	InApp     bool
	InContact bool
}

func (pd *PageData) SetFeaturesFromConfig() {
	pd.Feature = map[string]bool{
		"Blog":    ServerConfig.Features.Blog,
		"App":     ServerConfig.Features.App,
		"Contact": ServerConfig.Features.Contact,
		"Twitter": ServerConfig.Features.Twitter,
		"GitHub":  ServerConfig.Features.GitHub,
		"Analytics":  ServerConfig.Features.Analytics,
		"Shariff":  ServerConfig.Features.Shariff,
	}
	pd.CopyrightLine = ServerConfig.StaticTexts.CopyrightLine
	pd.InApp = false
}
