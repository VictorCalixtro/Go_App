package main

import (
	"encoding/gob"
	"github.com/alexedwards/scs/v2"
	"github.com/victorcalixtro/Web_App/internal/config"
	"github.com/victorcalixtro/Web_App/internal/handlers"
	"github.com/victorcalixtro/Web_App/internal/helpers"
	"github.com/victorcalixtro/Web_App/internal/models"
	"github.com/victorcalixtro/Web_App/internal/render"
	"log"
	"net/http"
	"os"
	"time"
)

const portnumber = ":8080"
var app config.AppConfig
var session *scs.SessionManager
var infoLog *log.Logger
var errorLog *log.Logger

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
	infoLog = log.New(os.Stdout, "INFO\t",log.Ldate | log.Ltime)
	app.InfoLog = infoLog

	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate | log.Ltime | log.Lshortfile)
	app.ErrorLog = errorLog


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
	helpers.NewHelpers(&app)


	log.Println("Serving on port", portnumber)

	return nil
}