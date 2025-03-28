package data

import "time"

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

// GetCustomerGrid returns all customers
func (c *Customer) GetCustomerGrid(pageNum uint) ([]*Customer, error) {
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
