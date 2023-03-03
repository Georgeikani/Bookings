package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/georgeikani/Bookings/internal/config"
	"github.com/georgeikani/Bookings/internal/models"
	"github.com/georgeikani/Bookings/render"
)

var Repo *Repository

// Repository is the respository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, r, "home.html", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// perform some logic

	StringMap := make(map[string]string)
	StringMap["test"] = "Hello, how are you?"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	StringMap["remote_ip"] = remoteIP

	//send data to template
	render.RenderTemplate(w, r, "about.html", &models.TemplateData{
		StringMap: StringMap,
	})
}

// General renders the general page
func (m *Repository) General(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "general.html", &models.TemplateData{})
}

// General renders the majors page
func (m *Repository) Major(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "major.html", &models.TemplateData{})
}

// Reservation renders the reservation page and displays form
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "reservation.html", &models.TemplateData{})
}

// PostReservation renders the reservation page and displays form
func (m *Repository) PostReservation(w http.ResponseWriter, r *http.Request) {
	start := r.Form.Get("start")
	end := r.Form.Get("end")

	w.Write([]byte(fmt.Sprintf("start date is %s and end is %s", start, end)))
}

type jsonResponse struct {
	Ok bool `json: "ok"`
	Message string `json:"message"`

}

// ReservationJSON handles request reservation and send JSON response
func (m *Repository) ReservationJSON(w http.ResponseWriter, r *http.Request) {
resp := jsonResponse{
	Ok: false,
	Message: "Available",

}
out, err := json.MarshalIndent(resp, "", "     ")
if err != nil {
	log.Println(err)
}

log.Println(string(out))
w.Header().Set("Content-Type", "application/json")
w.Write(out)
}

// Search available renders the available rooms page and displays the form
func (m *Repository) Search(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "search-availability.html", &models.TemplateData{})
}

// Search available renders the available rooms page and displays the form
func (m *Repository) Reserve(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "make-reservation.html", &models.TemplateData{})
}

// Contacts renders the contact page and displays the form
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "contact.html", &models.TemplateData{})
}
