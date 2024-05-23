package main

import (
	"fmt"
	"task/api"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Starting server on localhost:8080")
	router := gin.Default()

	// Set up routes
	api.RegisterRoutes(router)

	router.Run("localhost:8080")
}
