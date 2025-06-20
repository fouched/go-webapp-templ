package repo

import (
	"github.com/fouched/go-webapp-templ/internal/models"
)

type testCustomerRepo struct {
}

func (r *testCustomerRepo) SelectCustomerGrid(page int) ([]models.Customer, error) {
	return nil, nil
}

func (r *testCustomerRepo) SelectCustomerGridWithFilter(page int, filter string) ([]models.Customer, error) {
	return nil, nil
}

func (r *testCustomerRepo) SelectCustomerById(id int64) (models.Customer, error) {
	return models.Customer{}, nil
}

func (r *testCustomerRepo) CustomerInsert(customer *models.Customer) (int64, error) {
	return 0, nil
}

func (r *testCustomerRepo) CustomerUpdate(customer *models.Customer) error {
	return nil
}

func (r *testCustomerRepo) CustomerDelete(id int64) error {
	return nil
}
