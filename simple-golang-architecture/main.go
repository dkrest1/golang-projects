package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("hello world")
	// Create a new Gin router
	r := gin.Default()

	// Define a route
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, Gin!",
		})
	})

	// Run the server on port 8080
	r.Run(":8080")
}