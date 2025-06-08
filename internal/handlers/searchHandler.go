package handlers

import (
	"github.com/fouched/go-webapp-templ/internal/render"
	"github.com/fouched/go-webapp-templ/internal/services"
	"github.com/fouched/go-webapp-templ/internal/templates"
	"net/http"
	"strconv"
	"strings"
)

func (h *Handlers) Search(w http.ResponseWriter, r *http.Request) {
	page, ok := h.App.Session.Get(r.Context(), "page").(string)
	// this should never happen...
	if !ok {
		h.App.ErrorLog.Println("No session data for search exiting ")
		return
	}

	pageNum := 0
	filter := strings.TrimLeft(r.URL.Query().Get("filter"), " ")

	if page == "customer" {
		customerService := services.NewCustomerService(h.App)
		customers, err := customerService.GetCustomerGrid(pageNum, filter)
		if err != nil {
			h.App.ErrorLog.Print(err)
			return
		}

		t := templates.CustomerSearch(customers, strconv.Itoa(pageNum), filter)
		_ = render.Template(w, r, t)
	}
}
