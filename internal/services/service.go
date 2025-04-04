package services

import (
	"github.com/fouched/go-webapp-templ/internal/data"
	"github.com/fouched/go-webapp-templ/internal/models"
)

type CustomerServicer interface {
	GetCustomerGrid(page int, filter string) ([]models.Customer, error)
	GetCustomerGridV2(page uint, filter string) ([]*data.Customer, error)
	GetCustomerById(id int64) (models.Customer, error)
	GetCustomerByIdV2(id int64) (*data.Customer, error)
	CustomerInsert(customer *models.Customer) (int64, error)
	CustomerInsertV2(customer *data.Customer) (int64, error)
	CustomerUpdate(customer *models.Customer) error
	CustomerUpdateV2(customer *data.Customer) error
	DeleteCustomerById(id int64) error
	DeleteCustomerByIdV2(id int64) error
}
