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
		customers, err := services.CustomerService(h.App).GetCustomerGrid(pageNum, filter)
		if err != nil {
			h.App.ErrorLog.Print(err)
			return
		}

		t := templates.CustomerSearch(customers, strconv.Itoa(pageNum), filter)
		_ = render.Template(w, r, t)
	}

}

func (h *Handlers) SearchV2(w http.ResponseWriter, r *http.Request) {

	page, ok := h.App.Session.Get(r.Context(), "page").(string)
	// this should never happen...
	if !ok {
		h.App.ErrorLog.Println("No session data for search exiting ")
		return
	}

	pageNum := 1
	filter := strings.TrimLeft(r.URL.Query().Get("filter"), " ")

	if page == "customer" {
		customers, err := services.CustomerService(h.App).GetCustomerGridV2(uint(pageNum), filter)
		if err != nil {
			h.App.ErrorLog.Print(err)
			return
		}

		// upper starts pagination on 1, increment for pagination
		t := templates.CustomerSearchV2(customers, strconv.Itoa(pageNum+1), filter)
		_ = render.Template(w, r, t)
	}

}
