package main

import (
	"main/controllers"
	"main/database"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	database, err := database.Connect()
	if err != nil {
		panic("Failed to connect to the database")
	}

	productController := controllers.CreateProductController(database)

	server.GET("/ping", controllers.PingHandler)
	server.GET("/products", productController.GetProductController)

	server.Run(":8000")
}
