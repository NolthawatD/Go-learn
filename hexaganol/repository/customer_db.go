package repository

import "github.com/jmoiron/sqlx"

// Adapter

type customerRepository struct {
	db *sqlx.DB
}

func NewCustomerRepositoryDB(db *sqlx.DB) customerRepository {
	return customerRepository{db: db}
}

func (r customerRepository) GetAll() ([]Customer, error) {
	customers := []Customer{}
	query := "select customer_id, name, date_of_birth, city, zipcode, status from customers"
	err := r.db.Select(&customers, query)
	if err != nil {
		return nil, err
	}
	return customers, nil
}

func (r customerRepository) GetById(id int) (*Customer, error) {
	customer := Customer{}
	query := "select customer_id, name, date_of_birth, city, zipcode, status from customers where customer_id=?"
	err := r.db.Get(&customer, query, id)
	if err != nil {
		return nil, err
	}
	return &customer, nil
}
