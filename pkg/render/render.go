package render

import (
	"bytes"
	"fmt"
	"github.com/victorcalixtro/Web_App/pkg/config"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{}

var app *config.AppConfig
//sets the config for the template package

func NewTemplates(a *config.AppConfig){
	app = a
}

// RenderTemplate renders a template
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	var tc map[string]*template.Template
	//get the template cache from the app config
	if app.UseCache {
		tc = app.TemplateCache
	} else{
		tc, _ = CreateTemplateCache()
	}

	t, ok := tc[tmpl]
	if !ok{
		log.Fatal("could not get template from temp cache")
	}

	buf := new(bytes.Buffer)

	_ = t.Execute(buf,nil)
	buf.WriteTo(w)


	_,err := buf.WriteTo(w)

	if err != nil {
		fmt.Println("Error writing template to browser",err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {

	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		fmt.Println("Page is currently", page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}

	return myCache, nil
}