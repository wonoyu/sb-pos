package router

import (
	"sb-pos/constants"
	"sb-pos/controllers"
	"sb-pos/middlewares"

	"github.com/gin-gonic/gin"
)

func initCouponRoute(router *gin.Engine) {
	router.POST(constants.Coupons, middlewares.JWTAuthAdmin(), controllers.InsertCoupon)
	router.GET(constants.Coupons, middlewares.JWTAuth(), controllers.GetCoupons)
	router.GET(constants.CouponById, middlewares.JWTAuth(), controllers.GetCouponById)
	router.PUT(constants.CouponById, middlewares.JWTAuthAdmin(), controllers.UpdateCoupon)
	router.DELETE(constants.CouponById, middlewares.JWTAuthAdmin(), controllers.DeleteCoupon)
}
