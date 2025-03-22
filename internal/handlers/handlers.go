package handlers

import (
	"github.com/fouched/go-webapp-templ/internal/config"
	"github.com/fouched/go-webapp-templ/internal/render"
	"net/http"
)

// Instance the setup for all handlers
var Instance *Handlers

type Handlers struct {
	App *config.App
}

// NewHandlerConfig set the handler config and services
func NewHandlerConfig(a *config.App) *Handlers {
	return &Handlers{
		App: a,
	}
}

// NewHandlers creates the handler instance
func NewHandlers(h *Handlers) {
	Instance = h
}

func getNotifications(r *http.Request) *render.Notification {
	return &render.Notification{
		Success: Instance.App.Session.PopString(r.Context(), "success"),
		Warning: Instance.App.Session.PopString(r.Context(), "warning"),
		Error:   Instance.App.Session.PopString(r.Context(), "error"),
	}
}
