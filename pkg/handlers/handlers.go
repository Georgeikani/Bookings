package handlers

import (
	"net/http"

	"github.com/georgeikani/Bookings/pkg/models"
	"github.com/georgeikani/Bookings/pkg/config"
	"github.com/georgeikani/Bookings/render"
)

var Repo *Repository

//Repository is the respository type
type Repository struct {
	App *config.AppConfig
}
//NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

//NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
remoteIP := r.RemoteAddr
m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.html", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// perform some logic

	StringMap := make(map[string]string)
	StringMap["test"] = "Hello, how are you?"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	StringMap["remote_ip"] = remoteIP

	//send data to template
	render.RenderTemplate(w, "about.html", &models.TemplateData{
		StringMap: StringMap,
	})
}
func (m *Repository) Signup(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "signup.html", &models.TemplateData{})
}
