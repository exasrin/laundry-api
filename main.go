package main

import (
	"database/sql"
	"fmt"
	"go-api-enigma/config"
	_ "github.com/lib/pq"
)

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
	if error != nill {
		fmt.Println(error)
	}
	db.Exec("INSERT INTO")
}
