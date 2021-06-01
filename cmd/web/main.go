package main

import (
	"encoding/gob"
	"github.com/alexedwards/scs/v2"
	"github.com/victorcalixtro/Web_App/internal/config"
	"github.com/victorcalixtro/Web_App/internal/handlers"
	"github.com/victorcalixtro/Web_App/internal/models"
	"github.com/victorcalixtro/Web_App/internal/render"
	"log"
	"net/http"
	"time"
)

const portnumber = ":8080"
var app config.AppConfig
var session *scs.SessionManager


func main() {
	err := run()
	if err != nil{
		log.Fatal(err)
	}


	serve := &http.Server{
		Addr: portnumber,
		Handler: routes(&app),

	}

	err = serve.ListenAndServe()
	log.Fatal(err)


}


func run() error {

	gob.Register(models.Reservation{})

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
		return err
	}
	app.TemplateCache = tc
	app.UseCache = false
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)
	log.Println("Serving on port", portnumber)



	return nil
}