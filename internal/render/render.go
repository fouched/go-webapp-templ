package render

import (
	"github.com/a-h/templ"
	"net/http"
)

type Notification struct {
	Success string
	Warning string
	Error   string
}

func Template(w http.ResponseWriter, r *http.Request, template templ.Component) error {
	return template.Render(r.Context(), w)
}
