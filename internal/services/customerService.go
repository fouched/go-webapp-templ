package services

import (
	"github.com/fouched/go-webapp-templ/internal/config"
	"github.com/fouched/go-webapp-templ/internal/data"
	"github.com/fouched/go-webapp-templ/internal/models"
	"github.com/fouched/go-webapp-templ/internal/repo"
)

var customerService *customerServicer

type customerServicer struct {
	Repo repo.CustomerRepo
	App  *config.App
}

func CustomerService(a *config.App) CustomerServicer {

	if customerService == nil {
		customerService = &customerServicer{
			Repo: repo.NewCustomerRepo(a),
			App:  a,
		}
	}

	return customerService
}

func (s *customerServicer) GetCustomerGrid(page int, filter string) ([]models.Customer, error) {

	if filter == "" {
		customers, err := s.Repo.SelectCustomerGrid(page)
		if err != nil {
			return nil, err
		}

		return customers, nil
	} else {
		customers, err := s.Repo.SelectCustomerGridWithFilter(page, filter)
		if err != nil {
			return nil, err
		}

		return customers, nil
	}

}

func (s *customerServicer) GetCustomerGridV2(page uint, filter string) ([]*data.Customer, error) {

	if filter == "" {
		customers, err := s.App.Repo.Customers.GetGrid(page)
		if err != nil {
			return nil, err
		}

		return customers, nil
	} else {
		customers, err := s.App.Repo.Customers.GetGrid(page)
		if err != nil {
			return nil, err
		}

		return customers, nil
	}
}

func (s *customerServicer) GetCustomerById(id int64) (models.Customer, error) {

	customer, err := s.Repo.SelectCustomerById(id)
	if err != nil {
		return models.Customer{}, err
	}

	return customer, nil
}

func (s *customerServicer) GetCustomerByIdV2(id int64) (*data.Customer, error) {

	customer, err := s.App.Repo.Customers.Get(id)
	if err != nil {
		return nil, err
	}

	return customer, nil
}

func (s *customerServicer) CustomerInsert(customer *models.Customer) (int64, error) {

	id, err := s.Repo.CustomerInsert(customer)
	return id, err
}

func (s *customerServicer) CustomerInsertV2(customer *data.Customer) (int64, error) {

	id, err := s.App.Repo.Customers.Add(customer)
	return id, err
}

func (s *customerServicer) CustomerUpdate(customer *models.Customer) error {

	return s.Repo.CustomerUpdate(customer)
}

func (s *customerServicer) CustomerUpdateV2(customer *data.Customer) error {

	return s.App.Repo.Customers.Update(customer)
}

func (s *customerServicer) DeleteCustomerById(id int64) error {

	return s.Repo.CustomerDelete(id)
}

func (s *customerServicer) DeleteCustomerByIdV2(id int64) error {

	return s.App.Repo.Customers.Delete(id)
}
