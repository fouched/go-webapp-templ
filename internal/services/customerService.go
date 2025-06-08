package services

import (
	"github.com/fouched/go-webapp-templ/internal/config"
	"github.com/fouched/go-webapp-templ/internal/models"
	"github.com/fouched/go-webapp-templ/internal/repo"
)

type CustomerService struct {
	Repo repo.CustomerRepoInterface
	App  *config.App
}

func NewCustomerService(app *config.App, repo repo.CustomerRepoInterface) CustomerServiceInterface {
	return &CustomerService{
		Repo: repo,
		App:  app,
	}
}

func (s *CustomerService) GetCustomerGrid(page int, filter string) ([]models.Customer, error) {
	// overwrite the default repo to use the upper/db one
	s.Repo = repo.NewCustomerRepoUpperDB(s.App.DB)

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

func (s *CustomerService) GetCustomerById(id int64) (models.Customer, error) {
	customer, err := s.Repo.SelectCustomerById(id)
	if err != nil {
		return models.Customer{}, err
	}

	return customer, nil
}

func (s *CustomerService) CustomerInsert(customer *models.Customer) (int64, error) {
	id, err := s.Repo.CustomerInsert(customer)
	return id, err
}

func (s *CustomerService) CustomerUpdate(customer *models.Customer) error {
	return s.Repo.CustomerUpdate(customer)
}

func (s *CustomerService) DeleteCustomerById(id int64) error {
	return s.Repo.CustomerDelete(id)
}
