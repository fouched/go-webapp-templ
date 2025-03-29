package data

import (
	"time"
)

type Customer struct {
	ID           int64     `db:"id,omitempty"`
	CustomerName string    `db:"customer_name"`
	Tel          string    `db:"tel"`
	Email        string    `db:"email"`
	Address1     string    `db:"address_1"`
	Address2     string    `db:"address_2"`
	Address3     string    `db:"address_3"`
	PostCode     string    `db:"post_code"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}

func (c *Customer) Table() string {

	return "customer"
}

// GetGrid returns all customers
func (c *Customer) GetGrid(pageNum uint) ([]*Customer, error) {

	var customers []*Customer

	collection := upper.Collection(c.Table())
	rs := collection.Find().OrderBy("customer_name")
	p := rs.Paginate(PageSize)

	err := p.Page(pageNum).All(&customers)
	if err != nil {
		return nil, err
	}

	return customers, nil
}

// GetGridFiltered returns a filtered slice of customers based on criteria
func (c *Customer) GetGridFiltered(pageNum uint, filter string) ([]*Customer, error) {

	var customers []*Customer

	rs := upper.SQL().SelectFrom(c.Table()).
		Where("customer_name ILIKE ?", "%"+filter+"%").
		OrderBy("customer_name")

	p := rs.Paginate(PageSize)
	err := p.Page(pageNum).All(&customers)
	if err != nil {
		return nil, err
	}

	return customers, nil
}

func (c *Customer) Get(id int64) (*Customer, error) {

	var customer Customer

	collection := upper.Collection(c.Table())
	rs := collection.Find(id) // same as Find(upperdb.Cond{"id": id})
	err := rs.One(&customer)
	if err != nil {
		return &Customer{}, nil
	}

	return &customer, nil
}

func (c *Customer) Add(customer *Customer) (int64, error) {

	collection := upper.Collection(c.Table())
	rs, err := collection.Insert(customer)
	if err != nil {
		return 0, err
	}

	id := rs.ID()
	return id.(int64), nil
}

func (c *Customer) Update(customer *Customer) error {

	customer.UpdatedAt = time.Now()
	collection := upper.Collection(c.Table())
	rs := collection.Find(customer.ID)

	err := rs.Update(&customer)
	if err != nil {
		return err
	}

	return nil
}

func (c *Customer) Delete(id int64) error {

	collection := upper.Collection(c.Table())
	rs := collection.Find(id)

	err := rs.Delete()
	if err != nil {
		return err
	}

	return nil
}
