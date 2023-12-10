package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Login struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func main() {
	r := gin.Default()

	r.GET("/auth", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"response": "hello world"})
	})

	r.POST("/login", func(c *gin.Context) {
		var login Login

		if err := c.ShouldBindJSON(&login); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if login.Username != "admin" || login.Password != "admin" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})

	r.Run(":8080")
}
