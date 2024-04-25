package router

import (
	"sb-pos/constants"
	"sb-pos/controllers"
	"sb-pos/middlewares"

	"github.com/gin-gonic/gin"
)

func initOrderStatusRoute(router *gin.Engine) {
	router.POST(constants.OrderStatus, middlewares.JWTAuthAdmin(), controllers.InsertOrderStatus)
	router.GET(constants.OrderStatus, middlewares.JWTAuthAdmin(), controllers.GetOrderStatus)
	router.GET(constants.OrderStatusById, middlewares.JWTAuthAdmin(), controllers.GetOrderStatusById)
	router.PUT(constants.OrderStatusById, middlewares.JWTAuthAdmin(), controllers.UpdateOrderStatus)
	router.DELETE(constants.OrderStatusById, middlewares.JWTAuthAdmin(), controllers.DeleteOrderStatus)
}

func initOrdersRoute(router *gin.Engine) {
	router.GET(constants.Orders, middlewares.JWTAuth(), controllers.GetOrders)
	router.POST(constants.Orders, middlewares.JWATAuthCustomer(), controllers.CreateOrder)
	router.POST(constants.OrdersPay, middlewares.JWATAuthCustomer(), controllers.PayOrder)
	router.POST(constants.OrdersComplete, middlewares.JWATAuthCashier(), controllers.CompleteOrder)
}
