package services

import (
	"github.com/fouched/go-webapp-templ/internal/models"
)

type CustomerServiceInterface interface {
	GetCustomerGrid(page int, filter string) ([]models.Customer, error)
	GetCustomerById(id int64) (models.Customer, error)
	CustomerInsert(customer *models.Customer) (int64, error)
	CustomerUpdate(customer *models.Customer) error
	DeleteCustomerById(id int64) error
}
