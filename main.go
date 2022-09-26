package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type item struct {
	ID       string
	Name     string
	Value    int
	Quantity int
}

var items = []item{
	{ID: "1", Name: "VOLKSWAGEN POLO 6", Value: 15800, Quantity: 1},
	{ID: "2", Name: "Iphone 14", Value: 1000, Quantity: 1},
}

func getItems(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, items)
}

func main() {
	router := gin.Default()
	router.GET("/items", getItems)
	router.Run("localhost:9090")
}
