package router

import (
	"sb-pos/constants"
	"sb-pos/controllers"
	"sb-pos/middlewares"

	"github.com/gin-gonic/gin"
)

func initCustomerRouter(router *gin.Engine) {
	router.GET(constants.CustomerProfile, middlewares.JWATAuthCustomer(), controllers.GetProfile)
	router.POST(constants.CustomerTopup, middlewares.JWATAuthCustomer(), controllers.TopUp)
}
