package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	db "golang_net_worth_calculator_api/database"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db.Open()
	sqlDB, _ := db.DB.DB()

	defer sqlDB.Close()

	router := gin.Default()
	router.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "hello world!"})
	})

	port := os.Getenv("PORT")

	if port == "" {
		port = "9090"
	}
	// routes.AuthRoutes(router)
	// routes.UserRoutes(router)
	router.Run("localhost:9090")
	fmt.Println("Established a successful connection!")
}
