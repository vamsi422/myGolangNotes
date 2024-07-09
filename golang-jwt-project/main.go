package main

import (
	"fmt"
	"os"

	routes "./routes/"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("welcome to golang-jwt-project.")

	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())

	routes.AuthRouters(router)
	routes.UserRoutes(router)

	router.GET("/API-1", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"success": "Access granted for API-1"})
	})
	router.GET("/API-2", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"success": "Access granted for API-2"})
	})
	router.Run(":" + port)
}
