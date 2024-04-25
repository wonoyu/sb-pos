package router

import (
	"sb-pos/constants"
	"sb-pos/controllers"
	"sb-pos/middlewares"

	"github.com/gin-gonic/gin"
)

func initUserRoute(router *gin.Engine) {
	router.GET(constants.Users, middlewares.JWTAuth(), controllers.GetAllUsers)
	router.GET(constants.UserById, middlewares.JWTAuth(), controllers.GetUserById)
	router.PUT(constants.UserById, middlewares.JWTAuth(), controllers.UpdateUser)
	router.PUT(constants.UpdateUserRole, middlewares.JWTAuthAdmin(), controllers.UpdateUserRole)
	router.DELETE(constants.UserById, middlewares.JWTAuthAdmin(), controllers.DeleteUser)
}

func initAuthRoute(router *gin.Engine) {
	router.POST(constants.Login, controllers.Login)
	router.POST(constants.RegisterAdmin, controllers.RegisterAdmin)
	router.POST(constants.RegisterCustomer, controllers.RegisterCustomer)
	router.POST(constants.RegisterCashier, controllers.RegisterCashier)
}
