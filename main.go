package main

import (
	"github.com/gin-gonic/gin"
)

type item struct {
	Id       string
	name     string
	value    string
	quantity int
}

func main() {
	router := gin.Default()
	router.Run("localhost:9090")
}
