package routers

import (
	"database/sql"
	"main/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(server *gin.Engine, dbConnection *sql.DB) {

	userController := controllers.CreateUserController(dbConnection)
	server.GET("/users", userController.GetUserController)
}
