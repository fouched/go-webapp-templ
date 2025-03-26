package data

import (
	"database/sql"
	upperdb "github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/postgresql"
)

const PageSize = 20

var db *sql.DB
var upper upperdb.Session

type Models struct {
	Customers Customer
}

func New(databasePool *sql.DB) Models {
	db = databasePool
	upper, _ = postgresql.New(db)

	return Models{
		Customers: Customer{},
	}
}
