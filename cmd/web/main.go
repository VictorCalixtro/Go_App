package main

import (
	"github.com/victorcalixtro/Web_App/pkg/config"
	"github.com/victorcalixtro/Web_App/pkg/handlers"
	"github.com/victorcalixtro/Web_App/pkg/render"
	"log"
	"net/http"
)

const portnumber = ":8080"


func main() {
	var app config.AppConfig
	tc, err := render.CreateTemplateCache()
	if err != nil{
		log.Fatal("cannot load template cache")
	}
	app.TemplateCache = tc
	app.UseCache = false
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)
	log.Println("Serving on port", portnumber)
	serve := &http.Server{
		Addr: portnumber,
		Handler: routes(&app),

	}

	err = serve.ListenAndServe()
	log.Fatal(err)


}