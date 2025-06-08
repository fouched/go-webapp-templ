package services

import (
	"github.com/fouched/go-webapp-templ/internal/config"
	"github.com/fouched/go-webapp-templ/internal/models"
	"github.com/fouched/go-webapp-templ/internal/repo"
)

type CustomerService struct {
	App          *config.App
	CustomerRepo repo.CustomerRepoInterface
}

func NewCustomerService(app *config.App) CustomerServiceInterface {
	return &CustomerService{
		App:          app,
		CustomerRepo: app.Repo.CustomerRepo, // just a shortcut to make service funcs easier to work with
	}
}

func (s *CustomerService) GetCustomerGrid(page int, filter string) ([]models.Customer, error) {

	if filter == "" {
		customers, err := s.CustomerRepo.SelectCustomerGrid(page)
		if err != nil {
			return nil, err
		}

		return customers, nil
	} else {
		customers, err := s.CustomerRepo.SelectCustomerGridWithFilter(page, filter)
		if err != nil {
			return nil, err
		}

		return customers, nil
	}
}

func (s *CustomerService) GetCustomerById(id int64) (models.Customer, error) {
	customer, err := s.CustomerRepo.SelectCustomerById(id)
	if err != nil {
		return models.Customer{}, err
	}

	return customer, nil
}

func (s *CustomerService) CustomerInsert(customer *models.Customer) (int64, error) {
	id, err := s.CustomerRepo.CustomerInsert(customer)
	return id, err
}

func (s *CustomerService) CustomerUpdate(customer *models.Customer) error {
	return s.CustomerRepo.CustomerUpdate(customer)
}

func (s *CustomerService) DeleteCustomerById(id int64) error {
	return s.CustomerRepo.CustomerDelete(id)
}
