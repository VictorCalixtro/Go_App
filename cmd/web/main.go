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
	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	log.Println("Serving on port", portnumber)
	_ =http.ListenAndServe(portnumber,nil)

}