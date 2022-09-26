package routes

import (
	controller "golang_net_worth_calculator_api/controllers"
	"golang_net_worth_calculator_api/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("users/", controller.GetUsers())
	incomingRoutes.GET("users/:_id", controller.GetUser())
}
