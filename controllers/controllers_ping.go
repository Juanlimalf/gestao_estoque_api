package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PingHandler(c *gin.Context) {

	fmt.Println("Ping received")

	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
