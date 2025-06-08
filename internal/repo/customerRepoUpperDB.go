package repo

import (
	"database/sql"
	"github.com/fouched/go-webapp-templ/internal/models"
	"github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/postgresql"
	"time"
)

type CustomerRepoV2 struct {
	Session db.Session
}

// NewCustomerRepoUpperDB initializes and returns a repository instance
func NewCustomerRepoUpperDB(db *sql.DB) CustomerRepoInterface {
	session, _ := postgresql.New(db)
	return &CustomerRepoV2{
		Session: session,
	}
}

func (c *CustomerRepoV2) Table() string {
	return "customer"
}

// SelectCustomerGrid returns all customers
func (c *CustomerRepoV2) SelectCustomerGrid(pageNum int) ([]models.Customer, error) {
	var customers []models.Customer

	collection := c.Session.Collection(c.Table())
	rs := collection.Find().OrderBy("customer_name")
	p := rs.Paginate(PageSize)

	err := p.Page(uint(pageNum)).All(&customers)
	if err != nil {
		return nil, err
	}

	return customers, nil
}

// SelectCustomerGridWithFilter returns a filtered slice of customers based on criteria
func (c *CustomerRepoV2) SelectCustomerGridWithFilter(page int, filter string) ([]models.Customer, error) {
	var customers []models.Customer

	rs := c.Session.SQL().SelectFrom(c.Table()).
		Where("customer_name ILIKE ?", "%"+filter+"%").
		OrderBy("customer_name")

	p := rs.Paginate(PageSize)
	err := p.Page(uint(page)).All(&customers)
	if err != nil {
		return nil, err
	}

	return customers, nil
}

func (c *CustomerRepoV2) SelectCustomerById(id int64) (models.Customer, error) {
	var customer models.Customer

	collection := c.Session.Collection(c.Table())
	rs := collection.Find(id) // same as Find(upperdb.Cond{"id": id})
	err := rs.One(&customer)
	if err != nil {
		return models.Customer{}, nil
	}

	return customer, nil
}

func (c *CustomerRepoV2) CustomerInsert(customer *models.Customer) (int64, error) {
	collection := c.Session.Collection(c.Table())
	rs, err := collection.Insert(customer)
	if err != nil {
		return 0, err
	}

	id := rs.ID()
	return id.(int64), nil
}

func (c *CustomerRepoV2) CustomerUpdate(customer *models.Customer) error {
	customer.UpdatedAt = time.Now()
	collection := c.Session.Collection(c.Table())
	rs := collection.Find(customer.ID)

	err := rs.Update(&customer)
	if err != nil {
		return err
	}

	return nil
}

func (c *CustomerRepoV2) CustomerDelete(id int64) error {
	collection := c.Session.Collection(c.Table())
	rs := collection.Find(id)

	err := rs.Delete()
	if err != nil {
		return err
	}

	return nil
}
