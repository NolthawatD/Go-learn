package main

import (
	"fmt"
	"hexagonal/repository"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func main() {

	db, err := sqlx.Open("mysql", "root:secret@tcp(127.0.0.1:3306)/bank?parseTime=true")
	if err != nil {
		panic(err)
	}

	customerRepository := repository.NewCustomerRepositoryDB((db))

	_ = customerRepository

	// customer, err := customerRepository.GetById(2000)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(customer)

	customers, err := customerRepository.GetAll()
	if err != nil {
		panic(err)
	}

	fmt.Println(customers)
}
