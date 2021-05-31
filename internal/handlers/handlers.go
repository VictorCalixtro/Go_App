package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/victorcalixtro/Web_App/internal/config"
	"github.com/victorcalixtro/Web_App/internal/forms"
	"github.com/victorcalixtro/Web_App/internal/models"
	"github.com/victorcalixtro/Web_App/internal/render"
	"log"
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

	render.RenderTemplate(w, r,"home.page.tmpl", &models.TemplateData{})
}


//About is the about page handler
func (m *Repository)About(w http.ResponseWriter, r *http.Request){
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."

	remoteIp := m.App.Session.GetString(r.Context(),"remote_ip")

	stringMap["remote_ip"] = remoteIp

	//sends the data to the template
	render.RenderTemplate(w, r,"about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// Reservation renders the make a reservation page and displays form
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "make-reservation.page.tmpl", &models.TemplateData{
		Form: forms.New(nil),
	})
}

// PostReservation post reservation handles the posting of a reservation form
func (m *Repository) PostReservation(w http.ResponseWriter, r *http.Request) {

}


// Generals renders the room page
func (m *Repository) Generals(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r,"generals.page.tmpl", &models.TemplateData{})
}

// Majors renders the room page
func (m *Repository) Majors(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w,r, "majors.page.tmpl", &models.TemplateData{})
}

// Availability renders the search availability page
func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r,"search-availability.page.tmpl", &models.TemplateData{})
}

// PostAvailability renders the search availability page
func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	start:=r.Form.Get("start")
	end:= r.Form.Get("end")


	w.Write([]byte(fmt.Sprintf("start data is %s and end date is %s", start ,end)))
}

type jsonResponse struct{
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}


// AvailabilityJson handles request for availability and sends json responce
func (m *Repository) AvailabilityJson(w http.ResponseWriter, r *http.Request) {
	resp:= jsonResponse{
		OK: true,
		Message: "Available",
	}

	out, err := json.MarshalIndent(resp,"","     ")
	if err != nil{
		log.Println(err)
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)

}

// Contact renders the contact page
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w,r, "contact.page.tmpl", &models.TemplateData{})
}
