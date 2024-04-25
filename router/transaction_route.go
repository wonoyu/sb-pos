package router

import (
	"sb-pos/constants"
	"sb-pos/controllers"
	"sb-pos/middlewares"

	"github.com/gin-gonic/gin"
)

func initTransactionTypeRoute(router *gin.Engine) {
	router.POST(constants.TransactionType, middlewares.JWTAuthAdmin(), controllers.InsertTransactionType)
	router.GET(constants.TransactionType, middlewares.JWTAuthAdmin(), controllers.GetTransactionType)
	router.GET(constants.TransactionTypeById, middlewares.JWTAuthAdmin(), controllers.GetTransactionTypeById)
	router.PUT(constants.TransactionTypeById, middlewares.JWTAuthAdmin(), controllers.UpdateTransactionType)
	router.DELETE(constants.TransactionTypeById, middlewares.JWTAuthAdmin(), controllers.DeleteTransactionType)
}

func initTransactionRoute(router *gin.Engine) {
	router.GET(constants.SalesTransactions, middlewares.JWATAuthCashier(), controllers.GetSalesTransactions)
	router.GET(constants.SalesTransactionById, middlewares.JWATAuthCashier(), controllers.GetSalesTransactionById)
}
