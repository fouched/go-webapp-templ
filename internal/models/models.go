package models

import "time"

type Customer struct {
	ID           int64     `db:"id,omitempty"`
	CustomerName string    `db:"customer_name" schema:"customerName"`
	Tel          string    `db:"tel" schema:"tel"`
	Email        string    `db:"email" schema:"email"`
	Address1     string    `db:"address_1" schema:"address1"`
	Address2     string    `db:"address_2" schema:"address2"`
	Address3     string    `db:"address_3" schema:"address3"`
	PostCode     string    `db:"post_code" schema:"postCode"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}
