package storage

import "database/sql"

type orderproductRepo struct {
	DB *sql.DB
}

func NewOrderProduct(db *sql.DB) orderproductRepo {
	return orderproductRepo{
		DB: db,
	}
}
