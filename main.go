package main

import (
	"go-api-enigma/delivery"
	// "go-api-enigma/model"
	// "net/http"

	// "github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Customer struct {
	Id          string
	Name        string
	PhoneNumber string
	Address     string
}

func main() {
	// routerEngine := gin.Default()

	// routerEngine.GET("/hello-world", func(ctx *gin.Context) {
	// 	ctx.String(http.StatusOK, "Hello World")
	// })

	// // param/ path variable
	// // http://localhost:8080/hello-world/asrin
	// routerEngine.GET("/hello-world/:id", func(ctx *gin.Context) {
	// 	id := ctx.Param("id")
	// 	ctx.String(http.StatusOK, "Berhasil mendapatkan id %s", id)
	// })

	// // query param
	// // http://localhost:8080/products?name=asrin
	// routerEngine.GET("/products", func(ctx *gin.Context) {
	// 	name := ctx.Query("name")
	// 	ctx.String(http.StatusOK, "Berhasil mencari product dengan nama %s", name)
	// })

	// routerEngine.GET("/products2/:id", func(ctx *gin.Context) {
	// 	id := ctx.Param("id")
	// 	product := model.Product{
	// 		Id:    id,
	// 		Name:  "setrika + cuci baju",
	// 		Price: 100000,
	// 		Uom: model.Uom{
	// 			Id:   "1",
	// 			Type: "kg",
	// 		},
	// 	}
	// 	ctx.JSON(http.StatusOK, product)
	// })

	// // POSY
	// routerEngine.POST("/products", func(ctx *gin.Context) {
	// 	var product model.Product
	// 	// shouldBind JSON -> struct
	// 	err := ctx.ShouldBindJSON(&product)
	// 	if err != nil {
	// 		ctx.JSON(http.StatusBadRequest, err.Error())
	// 		return
	// 	}
	// 	ctx.JSON(http.StatusCreated, product)
	// })

	// routerEngine.Run()

	delivery.NewServer().Run()
}
