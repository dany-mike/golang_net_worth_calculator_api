package controllers

import (
	"github.com/go-playground/validator/v10"
)

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
