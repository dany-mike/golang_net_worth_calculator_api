package controllers

import (
	"golang_net_worth_calculator_api/database"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")
var validate = validator.New()

func HashPassword() string {
	return "Hash password"
}

func VerifyPassword() string {
	return "Verify password"
}

func Signup() string {
	return "Signup"
}

func Login() string {
	return "Login"
}

func GetUsers() string {
	return "get users"
}

func GetUser() string {
	return "get user"
}
