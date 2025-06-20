package handlers

import (
	"fmt"
	"github.com/fouched/go-webapp-templ/internal/models"
	"github.com/fouched/go-webapp-templ/internal/render"
	"github.com/fouched/go-webapp-templ/internal/services"
	"github.com/fouched/go-webapp-templ/internal/templates"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func (h *Handlers) CustomerGrid(w http.ResponseWriter, r *http.Request) {
	page := "customer"
	h.App.Session.Put(r.Context(), "page", page)

	p := 0
	pageNum := r.URL.Query().Get("pageNum")
	filter := r.URL.Query().Get("filter")

	if pageNum != "" {
		p, _ = strconv.Atoi(pageNum)
	}

	customerService := services.NewCustomerService(h.App)
	customers, err := customerService.GetCustomerGrid(p, filter)
	if err != nil {
		h.App.ErrorLog.Print(err)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// increment p for next page
	p = p + 1
	t := templates.CustomerGrid(customers, strconv.Itoa(p), filter, h.getNotifications(r))
	_ = render.Template(w, r, t)
}

func (h *Handlers) CustomerDetails(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	customerService := services.NewCustomerService(h.App)
	customer, err := customerService.GetCustomerById(id)
	if err != nil {
		h.App.ErrorLog.Print(err)
		h.App.Session.Put(r.Context(), "error", "Error getting customer")
	}

	t := templates.CustomerDetails(customer)
	_ = render.Template(w, r, t)
}

func (h *Handlers) CustomerAddGet(w http.ResponseWriter, r *http.Request) {
	t := templates.CustomerAdd()
	_ = render.Template(w, r, t)
}

func (h *Handlers) CustomerAddPost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		h.App.ErrorLog.Print(err)
		h.App.Session.Put(r.Context(), "error", "Error parsing data for customer insert")
	}

	var customer models.Customer
	if err := h.Decode(&customer, r.PostForm); err != nil {
		h.App.ErrorLog.Print(err)
		h.App.Session.Put(r.Context(), "error", "Error parsing data for customer insert")
	}

	customerService := services.NewCustomerService(h.App)
	id, err := customerService.CustomerInsert(&customer)
	if err != nil {
		h.App.ErrorLog.Print(err)
		if strings.Contains(err.Error(), "duplicate") {
			h.App.Session.Put(r.Context(), "error", "Error customer already exists")
		} else {
			h.App.Session.Put(r.Context(), "error", "Error inserting customer")
		}
	} else {
		// in a real app we might to do something with an inserted record
		h.App.InfoLog.Println(fmt.Sprintf("Inserted customer with id: %d", id))
		h.App.Session.Put(r.Context(), "success", "Successfully added customer")
		customer.ID = id
	}

	http.Redirect(w, r, "/customer", http.StatusSeeOther)
}

func (h *Handlers) CustomerUpdate(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		h.App.ErrorLog.Print(err)
		h.App.Session.Put(r.Context(), "error", "Error updating customer")
	}

	id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	customer := models.Customer{
		ID:           id,
		CustomerName: r.Form.Get("customerName"),
		Tel:          r.Form.Get("tel"),
		Email:        r.Form.Get("email"),
		Address1:     r.Form.Get("address1"),
		Address2:     r.Form.Get("address2"),
		Address3:     r.Form.Get("address3"),
		PostCode:     r.Form.Get("postCode"),
		UpdatedAt:    time.Now(),
	}

	customerService := services.NewCustomerService(h.App)
	err = customerService.CustomerUpdate(&customer)
	if err != nil {
		h.App.ErrorLog.Print(err)
		h.App.Session.Put(r.Context(), "error", "Error updating customer")

	}

	h.App.Session.Put(r.Context(), "success", "Customer updated successfully")

	t := templates.CustomerUpdate(customer, h.getNotifications(r))
	_ = render.Template(w, r, t)
}

func (h *Handlers) CustomerDelete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	customerService := services.NewCustomerService(h.App)
	err := customerService.DeleteCustomerById(id)
	if err != nil {
		h.App.ErrorLog.Print(err)
		h.App.Session.Put(r.Context(), "error", "Error deleting customer")

		// load the customer again because we normally remove the row from the UI
		customer, _ := customerService.GetCustomerById(id)
		t := templates.CustomerUpdate(customer, h.getNotifications(r))
		_ = render.Template(w, r, t)
	} else {
		h.App.Session.Put(r.Context(), "success", "Customer deleted successfully")
		t := templates.CustomerDelete(h.getNotifications(r))
		_ = render.Template(w, r, t)
	}
}
