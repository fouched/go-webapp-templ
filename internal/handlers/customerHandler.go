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

	customers, err := services.CustomerService(h.App).GetCustomerGrid(p, filter)
	if err != nil {
		h.App.ErrorLog.Print(err)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// increment p for next page
	p = p + 1
	t := templates.CustomerGrid(customers, strconv.Itoa(p), filter, &render.Notification{
		Success: h.App.Session.PopString(r.Context(), "success"),
		Warning: h.App.Session.PopString(r.Context(), "warning"),
		Error:   h.App.Session.PopString(r.Context(), "error"),
	})
	_ = render.Template(w, r, t)
}

func (h *Handlers) CustomerDetails(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	customer, err := services.CustomerService(h.App).GetCustomerById(id)
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
		h.App.Session.Put(r.Context(), "error", "Error parsing form for customer insert")
	}

	customer := models.Customer{
		CustomerName: r.Form.Get("customerName"),
		Tel:          r.Form.Get("tel"),
		Email:        r.Form.Get("email"),
		Address1:     r.Form.Get("address1"),
		Address2:     r.Form.Get("address2"),
		Address3:     r.Form.Get("address3"),
		PostCode:     r.Form.Get("postCode"),
		UpdatedAt:    time.Now(),
	}

	id, err := services.CustomerService(h.App).CustomerInsert(&customer)
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

	err = services.CustomerService(h.App).CustomerUpdate(&customer)
	if err != nil {
		customer.Error = "Error updating customer"
		h.App.ErrorLog.Print(err)
		h.App.Session.Put(r.Context(), "error", customer.Error)

	}

	customer.Success = "Customer updated successfully"
	h.App.Session.Put(r.Context(), "success", customer.Success)

	t := templates.CustomerUpdate(customer, &render.Notification{
		Success: customer.Success,
	})
	_ = render.Template(w, r, t)
}

func (h *Handlers) CustomerDelete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	err := services.CustomerService(h.App).DeleteCustomerById(id)
	if err != nil {
		h.App.ErrorLog.Print(err)
		h.App.Session.Put(r.Context(), "error", "Error deleting customer")

		// load the customer again because we normally remove the row from the UI
		customer, _ := services.CustomerService(h.App).GetCustomerById(id)
		t := templates.CustomerUpdate(customer, &render.Notification{
			Error: "Error deleting customer",
		})
		_ = render.Template(w, r, t)
	} else {
		h.App.Session.Put(r.Context(), "success", "Customer deleted successfully")
		t := templates.CustomerDelete(&render.Notification{
			Success: "Customer deleted successfully",
		})
		_ = render.Template(w, r, t)
	}
}
