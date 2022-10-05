package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	db "golang_net_worth_calculator_api/database"
	routes "golang_net_worth_calculator_api/routes"
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

	routes.AuthRoutes(router)
	// routes.UserRoutes(router)
	router.Run("localhost:9090")
	fmt.Println("Established a successful connection!")
}
