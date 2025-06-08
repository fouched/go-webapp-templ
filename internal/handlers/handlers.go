package handlers

import (
	"github.com/fouched/go-webapp-templ/internal/config"
	"github.com/fouched/go-webapp-templ/internal/render"
	"net/http"
)

type Handlers struct {
	App *config.App
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
