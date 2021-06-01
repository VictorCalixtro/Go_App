package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/victorcalixtro/Web_App/internal/config"
	"github.com/victorcalixtro/Web_App/internal/forms"
	"github.com/victorcalixtro/Web_App/internal/helpers"
	"github.com/victorcalixtro/Web_App/internal/models"
	"github.com/victorcalixtro/Web_App/internal/render"
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
	render.RenderTemplate(w, r,"home.page.tmpl", &models.TemplateData{})
}


//About is the about page handler
func (m *Repository)About(w http.ResponseWriter, r *http.Request){
//sends the data to the template
	render.RenderTemplate(w, r,"about.page.tmpl", &models.TemplateData{})
}

// Reservation renders the make a reservation page and displays form
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	var emptyReservation models.Reservation
	data:= make(map[string]interface{})
	data["reservation"] = emptyReservation


	render.RenderTemplate(w, r, "make-reservation.page.tmpl", &models.TemplateData{
		Form: forms.New(nil),
	})
}

// PostReservation post reservation handles the posting of a reservation form
func (m *Repository) PostReservation(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		helpers.ServerError(w,err)
		return
	}

	reservation := models.Reservation{
		FirstName: r.Form.Get("first_name"),
		LastName: r.Form.Get("last_name"),
		Phone: r.Form.Get("phone"),
		Email: r.Form.Get("email"),
	}

	form := forms.New(r.PostForm)

	form.Required("first_name","last_name","email")
	form.MinLength("first_name", 3)
	form.IsEmail("email")

	if !form.Valid(){
		data := make(map[string]interface{})
		data["reservation"] = reservation

		render.RenderTemplate(w, r, "make-reservation.page.tmpl", &models.TemplateData{
			Form: form,
			Data: data,
		})
		return
	}

	m.App.Session.Put(r.Context(),"reservation",reservation)
	//Redirect
	http.Redirect(w,r,"/reservation-summary", http.StatusSeeOther)

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
		helpers.ServerError(w,err)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)

}

// Contact renders the contact page
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w,r, "contact.page.tmpl", &models.TemplateData{})
}

func (m *Repository) ReservationSummary(w http.ResponseWriter, r *http.Request){
	reservation, ok := m.App.Session.Get(r.Context(),"reservation").(models.Reservation)

	if !ok {
		m.App.ErrorLog.Println("Cant get error from session")
		m.App.Session.Put(r.Context(),"error","Cant get reservation from session")
		http.Redirect(w,r,"/",http.StatusTemporaryRedirect)
		return
	}

	data := make(map[string]interface{})
	data["reservation"] = reservation

	render.RenderTemplate(w,r, "reservation-summary.page.tmpl", &models.TemplateData{
		Data: data,
	})
}