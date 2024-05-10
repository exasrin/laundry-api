package main

import (
	"fmt"
	"go-api-enigma/config"
	"go-api-enigma/repository"

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

	conn, error := config.NewDbCOnnection(cfg)
	if error != nil {
		fmt.Println(error)
	}
	db := conn.Conn()

	// Customer := Customer{
	// 	Id:          "2",
	// 	Name:        "Rasyid Sulaiaman",
	// 	PhoneNumber: "081234567890",
	// 	Address:     "Kendari",
	// }

	// _, error = db.Exec("INSERT INTO m_customer VALUES ($1, $2, $3, $4)",
	// 	Customer.Id,
	// 	Customer.Name,
	// 	Customer.PhoneNumber,
	// 	Customer.Address,
	// )

	// if error != nil {
	// 	fmt.Println(error)
	// }

	// fmt.Println("success inserting data")

	// savee
	// uomRepo := repository.NewUomRepository(db)
	// uomRepo.Save(model.Uom{
	// 	Id:   "2",
	// 	Name: "ons",
	// })

	// getById
	uomRepo := repository.NewUomRepository(db)
	uom, err := uomRepo.FindById("1")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(uom)

	uoms, err := uomRepo.FindAll()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(uoms)
}
