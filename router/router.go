package router

import "github.com/gin-gonic/gin"

func InitRouter() (router *gin.Engine) {
	router = gin.Default()

	initUserRoute(router)
	initAuthRoute(router)
	initRoleRoute(router)
	initCategoryRoute(router)
	initProductRoute(router)
	// initCouponRoute(router)
	initOrderStatusRoute(router)
	initOrdersRoute(router)
	initCustomerRouter(router)
	initTransactionTypeRoute(router)
	initTransactionRoute(router)

	return
}
