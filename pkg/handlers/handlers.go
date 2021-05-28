package handlers

import (
	"github.com/victorcalixtro/Web_App/pkg/config"
	"github.com/victorcalixtro/Web_App/pkg/models"
	"github.com/victorcalixtro/Web_App/pkg/render"
	"net/http"
)




//Repo the repository used by the handlers
var Repo *Repository
//Repository is the repository type
type Repository struct{
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository{
	return &Repository{
		App: a,
	}
}

//NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository){
	Repo = r
}

//Home is the home page handler
func (m *Repository)Home(w http.ResponseWriter, r *http.Request){
	remoteIp := r.RemoteAddr //ipv4 or v6 address
	m.App.Session.Put(r.Context(),"remote_ip", remoteIp)

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}


//About is the about page handler
func (m *Repository)About(w http.ResponseWriter, r *http.Request){
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."

	remoteIp := m.App.Session.GetString(r.Context(),"remote_ip")

	stringMap["remote_ip"] = remoteIp

	//sends the data to the template
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}