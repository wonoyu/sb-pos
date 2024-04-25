package router

import (
	"sb-pos/constants"
	"sb-pos/controllers"
	"sb-pos/middlewares"

	"github.com/gin-gonic/gin"
)

func initProductRoute(router *gin.Engine) {
	router.POST(constants.Products, middlewares.JWTAuthAdmin(), controllers.InsertProduct)
	router.GET(constants.Products, middlewares.JWTAuth(), controllers.GetProducts)
	router.GET(constants.ProductById, middlewares.JWTAuth(), controllers.GetProductById)
	router.PUT(constants.ProductById, middlewares.JWTAuthAdmin(), controllers.UpdateProduct)
	router.DELETE(constants.ProductById, middlewares.JWTAuthAdmin(), controllers.DeleteProduct)
}

func initCategoryRoute(router *gin.Engine) {
	router.POST(constants.Categories, middlewares.JWTAuthAdmin(), controllers.InsertCategory)
	router.GET(constants.Categories, middlewares.JWTAuth(), controllers.GetCategories)
	router.GET(constants.CategoryById, middlewares.JWTAuth(), controllers.GetCategoryById)
	router.PUT(constants.CategoryById, middlewares.JWTAuthAdmin(), controllers.UpdateCategory)
	router.DELETE(constants.CategoryById, middlewares.JWTAuthAdmin(), controllers.DeleteCategory)
}
