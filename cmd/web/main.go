package main

import (
	"github.com/alexedwards/scs/v2"
	"github.com/victorcalixtro/Web_App/pkg/config"
	"github.com/victorcalixtro/Web_App/pkg/handlers"
	"github.com/victorcalixtro/Web_App/pkg/render"
	"log"
	"net/http"
	"time"
)

const portnumber = ":8080"
var app config.AppConfig
var session *scs.SessionManager


func main() {


	app.InProduction = false


	session =scs.New()
	session.Lifetime= 24 *time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction //true when in production since it will be https
	app.Session = session


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