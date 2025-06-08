package models

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
