package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/auth", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"response":"hello world"})
	})
	r.Run(":8080")
}
