package main

import (
	"hexagonal/handler"
	"hexagonal/repository"
	"hexagonal/service"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func main() {

	db, err := sqlx.Open("mysql", "root:secret@tcp(127.0.0.1:3306)/bank?parseTime=true")
	if err != nil {
		panic(err)
	}

	customerRepository := repository.NewCustomerRepositoryDB((db))
	customerService := service.NewCustomerService(customerRepository)
	customerHandler := handler.NewCustomerHandler(customerService)

	router := mux.NewRouter()

	router.HandleFunc("/customers", customerHandler.GetCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customerID:[0-9]+}", customerHandler.GetCustomer).Methods(http.MethodGet)

	http.ListenAndServe(":8000", router)
}
