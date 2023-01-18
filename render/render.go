package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/georgeikani/Bookings/pkg/models"
	"github.com/georgeikani/Bookings/pkg/config"
)

/*

// Building a simple template cache.
// This is how to render templates
//RenderTemplate renders template using html/templates

func RenderTemplates(w http.ResponseWriter, html string){
	parsedTemplate, _:= template.ParseFiles("./templates/" + html, "./templates/base.layout.html")
	err := parsedTemplate.Execute(w,nil)
	if err != nil {
		fmt.Println("error parsing template:", err)
		return
	}
}

var tc = make(map[string]*template.Template)

func RenderTemplateTest (w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error

	// check to see if we already have the template in our cache
	_, inMap := tc[t]
	if !inMap{
		// need to create the template
		log.Println("creating templates and adding to cache")
		err = createTemplateCache(t) // check for error
		if err != nil {
			log.Println(err)
		}
	}else{
		// we have the template in the cache
		log.Println("using cache template")
	}

	tmpl = tc[t]
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}
func createTemplateCache(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", t),
		"./templatates/base.layout.html",
	}

	//parse the template
tmpl, err := template.ParseFiles(templates...)
if err != nil {
	return err
}

tc[t] = tmpl
return nil
}

*/


var app *config.AppConfig

//NewTemplate set the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}


//Building more complex tempate cache
// RenderTemplate renders using html template
func RenderTemplate(w http.ResponseWriter, html string, td *models.TemplateData) {
	var tc map[string]*template.Template

	if app.UseCache {
		tc = app.TemplateCache
	}else {
		tc, _= CreateTemplateCache()
	}

	// get requested template from cache
	t, ok := tc[html]
	if !ok {
		log.Fatal("could not get template from template cache")
	}

	buf := new(bytes.Buffer)

	td = AddDefaultData(td)

	_ = t.Execute(buf, td)

	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("Error writing template to browser", err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	// get all of the files named *.page.tmpl from ./templates
	pages, err := filepath.Glob("./templates/*.html")
	if err != nil {
		return myCache, err
	}

	// rendering and ranging
	// range through all files ending with *.page.tmpl
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.html")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.html")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}

	return myCache, nil
}
