package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	"golang_net_worth_calculator_api/controllers"
)

var server = controllers.Server{}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	server.Initialize(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

	fmt.Println("Established a successful connection!")

	server.Run(":5050")
}
