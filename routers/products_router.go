package routers

import (
	"database/sql"
	"main/controllers"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(router *gin.Engine, dbConnection *sql.DB) {

	productController := controllers.CreateProductController(dbConnection)

	v1 := router.Group("/products")
	{
		v1.GET("/", productController.GetProductController)
		v1.POST("/", productController.CreateProductController)
	}

}
