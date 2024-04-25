package router

import (
	"sb-pos/constants"
	"sb-pos/controllers"
	"sb-pos/middlewares"

	"github.com/gin-gonic/gin"
)

func initRoleRoute(router *gin.Engine) {
	router.POST(constants.Roles, middlewares.JWTAuthAdmin(), controllers.InsertRole)
	router.GET(constants.Roles, middlewares.JWTAuthAdmin(), controllers.GetAllRoles)
	router.GET(constants.RoleById, middlewares.JWTAuthAdmin(), controllers.GetRoleById)
	router.PUT(constants.RoleById, middlewares.JWTAuthAdmin(), controllers.UpdateRole)
	router.DELETE(constants.RoleById, middlewares.JWTAuthAdmin(), controllers.DeleteRole)
}
