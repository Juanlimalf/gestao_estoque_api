package main

import (
	"main/database"
	"main/docs"
	"main/routers"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	basePath := "/"
	server := gin.Default()

	docs.SwaggerInfo.BasePath = basePath

	database, err := database.Connect()
	if err != nil {
		panic("Failed to connect to the database")
	}

	routers.ProductRoutes(server, database)
	routers.UserRoutes(server, database)

	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	server.Run(":8000")
}
