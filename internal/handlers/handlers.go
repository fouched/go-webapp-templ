package handlers

import (
	"github.com/fouched/go-webapp-templ/internal/config"
	"github.com/fouched/go-webapp-templ/internal/render"
	"github.com/gorilla/schema"
	"net/http"
)

type Handlers struct {
	App     *config.App
	Decoder *schema.Decoder
}

func NewHandlers(app *config.App) *Handlers {
	return &Handlers{App: app}
}

func (h *Handlers) getNotifications(r *http.Request) *render.Notification {
	return &render.Notification{
		Success: h.App.Session.PopString(r.Context(), "success"),
		Warning: h.App.Session.PopString(r.Context(), "warning"),
		Error:   h.App.Session.PopString(r.Context(), "error"),
	}
}

func (h *Handlers) Decode(dst interface{}, src map[string][]string) error {
	return h.Decoder.Decode(dst, src)
}
