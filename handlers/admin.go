// Copyright 2017 EcoSystem Software LLP

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

// 	http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package handlers

import (
	"fmt"
	"net/http"
	"path"
	"strings"

	eco "github.com/ecosystemsoftware/ecosystem/utilities"
	"github.com/spf13/afero"
	"github.com/spf13/viper"
	gin "gopkg.in/gin-gonic/gin.v1"
)

//AdminShowViews concatenates the contents of the views.json file in each bundle and returns the combined JSON
func AdminShowConcatenatedJSON(c *gin.Context) {

	//Work out whether this is views or menus etc
	url := c.Request.RequestURI

	urlParts := strings.Split(url, "/")
	stub := urlParts[1]

	//For each bundle present
	bundles := viper.GetStringSlice("bundlesInstalled")

	var compositeFileContents string

	for _, v := range bundles {

		//Work out the name of the views file
		viewsFile := path.Join("bundles", v, "admin-panel", stub+".json")

		//Check it exists
		ok, err := afero.Exists(eco.AppFs, viewsFile)
		//If it exists, try to read it
		if ok && err == nil {
			viewsFileContents, err := afero.ReadFile(eco.AppFs, viewsFile)
			//If it was read correctly
			if err == nil {

				//Prefix with the bundle name
				viewsFileString := fmt.Sprintf(`"%s":%s`, v, string(viewsFileContents))

				//If this is the first one, don't insert comma, otherwise do
				if compositeFileContents != "" {
					compositeFileContents += ","
				}
				compositeFileContents = fmt.Sprintf(`%s%s`, compositeFileContents, viewsFileString)
			}
		}

	}

	c.String(http.StatusOK, fmt.Sprintf(`{%s}`, compositeFileContents))

}

func AdminGetImports(c *gin.Context) {
	var cf map[string]string
	viper.Unmarshal(&cf)

	bundles := viper.GetStringSlice("bundlesInstalled")

	//TODO: template will try to import actions.html from each bundle, even if it doesn't exist
	c.HTML(http.StatusOK, "admin-imports.html", gin.H{
		"config":  cf,
		"bundles": bundles,
	})
}