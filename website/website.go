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

package website

import (
	"html/template"
	"log"

	"github.com/lib/pq"
)

type page struct {
	Records       []map[string]interface{}
	Schema, Table string
	Site          SiteBuilder
	HttpCode      int
	Message       string
	DBCode        pq.ErrorCode
}

//Shared templates holder
var templates *template.Template

//Activate is the main package activation function
func Activate() {
	parseTemplates()
	//Set the routes for the package
	setRoutes()
}

func parseTemplates() {
	log.Println("Parsing templates in bundles/**/templates/**/*.html ")
	//Start with the ecosystem.js templates
	//templates = template.Must(template.New("ecosystem.js").Parse(EcoSystemJS))

	templates = template.Must(template.New("defaulterror.html").Parse(DefaultErrorPage))
	//Add all the bundles templates
	templates, _ = templates.ParseGlob("bundles/**/templates/**/*.html")
}