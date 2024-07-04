package main

import (
	"fmt"
	"hexagonal/handler"
	"hexagonal/logs"
	"hexagonal/repository"
	"hexagonal/service"
	"net/http"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

func main() {
	initTimeZone()
	initConfit()
	db := initDatabase()

	customerRepositoryDB := repository.NewCustomerRepositoryDB((db))
	customerRepositoryMock := repository.NewCustomerRepositoryMock()
	_ = customerRepositoryMock

	customerService := service.NewCustomerService(customerRepositoryDB)
	customerHandler := handler.NewCustomerHandler(customerService)

	router := mux.NewRouter()

	router.HandleFunc("/customers", customerHandler.GetCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customerID:[0-9]+}", customerHandler.GetCustomer).Methods(http.MethodGet)

	port := fmt.Sprintf(":%v", viper.GetInt("app.port"))
	// log.Printf("Banking service started at port %v", port)
	logs.Info("Banking service started at port " + viper.GetString("app.port"))
	http.ListenAndServe(port, router)
}

func initConfit() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	// เวลา start project สามารถใส่ APP_PORT= 5000 go run .

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func initTimeZone() {
	ict, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		panic(err)
	}

	time.Local = ict
}

func initDatabase() *sqlx.DB {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true",
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.host"),
		viper.GetInt("db.port"),
		viper.GetString("db.database"),
	)

	db, err := sqlx.Open(viper.GetString("db.driver"), dsn)
	if err != nil {
		panic(err)
	}

	db.SetConnMaxIdleTime(3 * time.Minute)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db
}
