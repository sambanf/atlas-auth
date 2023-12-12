package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/steamyrain/atlas-auth/cmd/handlers"
)

func main() {
	r := gin.Default()

	r.GET("/auth", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"response": "hello world"})
	})

	r.POST("/login", handlers.HandleLogin)

	r.POST("/logout", handlers.HandleLogout)

	r.Run(":8080")
}
