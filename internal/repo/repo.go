package repo

import (
	"github.com/fouched/go-webapp-templ/internal/models"
	"time"
)

// DbTimeout should be lowered in production - 30 secs max
const DbTimeout = time.Minute * 2
const PageSize = 20

// CustomerRepoInterface defines repository methods
type CustomerRepoInterface interface {
	SelectCustomerGrid(page int) ([]models.Customer, error)
	SelectCustomerGridWithFilter(page int, filter string) ([]models.Customer, error)
	SelectCustomerById(int int64) (models.Customer, error)
	CustomerInsert(customer *models.Customer) (int64, error)
	CustomerUpdate(customer *models.Customer) error
	CustomerDelete(id int64) error
}
