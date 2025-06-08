package handlers

import (
	"github.com/fouched/go-webapp-templ/internal/render"
	"github.com/fouched/go-webapp-templ/internal/templates"
	"net/http"
)

func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	h.App.Session.Put(r.Context(), "searchTarget", "home")

	t := templates.Home()
	_ = render.Template(w, r, t)

}
