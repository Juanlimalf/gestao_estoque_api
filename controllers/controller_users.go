package controllers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	dbConnection *sql.DB
}

func CreateUserController(dbConnection *sql.DB) *UserController {
	return &UserController{dbConnection: dbConnection}
}

func (controller *UserController) GetUserController(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "GetUserController called"})
}
