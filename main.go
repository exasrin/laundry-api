package main

import (
	"database/sql"
	"fmt"
	"go-api-enigma/config"

	_ "github.com/lib/pq"
)

type Customer struct {
	Id          string
	Name        string
	PhoneNumber string
	Address     string
}

func main() {
	cfg, error := config.NewConfig()
	if error != nil {
		fmt.Println(error)
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DbConfig.Host,
		cfg.DbConfig.Port,
		cfg.DbConfig.User,
		cfg.DbConfig.Password,
		cfg.DbConfig.Name,
	)

	db, error := sql.Open(cfg.DbConfig.Driver, dsn)
	if error != nil {
		fmt.Println(error)
	}

	Customer := Customer{
		Id:          "1",
		Name:        "Alffathir Rasyid Sulaiaman",
		PhoneNumber: "081234567890",
		Address:     "Jakarta",
	}

	_, error = db.Exec("INSERT INTO m_customer VALUES ($1, $2, $3, $4)",
		Customer.Id,
		Customer.Name,
		Customer.PhoneNumber,
		Customer.Address,
	)

	if error != nil {
		fmt.Println(error)
	}

	fmt.Println("success inserting data")
}
