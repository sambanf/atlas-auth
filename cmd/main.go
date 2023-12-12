package main

import (
	"atlas-auth/cmd/handlers" // adjust this import path to match your project structure
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/auth", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"response": "hello world"})
	})

	r.POST("/login", handlers.Login)

	r.Run(":8080")
}
