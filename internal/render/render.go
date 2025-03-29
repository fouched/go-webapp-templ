package render

import (
	"github.com/a-h/templ"
	"github.com/fouched/go-webapp-templ/internal/config"
	"net/http"
)

var app *config.App

// NewRenderer sets the config for the template package
func NewRenderer(a *config.App) {
	app = a
}

type Notification struct {
	Success string
	Warning string
	Error   string
}

func Template(w http.ResponseWriter, r *http.Request, template templ.Component) error {

	// clear notification session data
	_ = app.Session.PopString(r.Context(), "success")
	_ = app.Session.PopString(r.Context(), "warning")
	_ = app.Session.PopString(r.Context(), "error")

	return template.Render(r.Context(), w)
}
