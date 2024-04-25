package controllers

import (
	"net/http"
	"sb-pos/database"
	"sb-pos/models"
	"sb-pos/repositories"
	"sb-pos/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetCoupons(c *gin.Context) {
	roles, err := repositories.GetCoupons(database.DbConnection)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": roles,
	})
}

func GetCouponById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	role, err := repositories.GetCouponById(database.DbConnection, id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": role,
	})
}

func InsertCoupon(c *gin.Context) {

	var coupon models.CreateCoupon

	err := utils.ValidateJsonFields(c, &coupon)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	err = repositories.InsertCoupon(database.DbConnection, coupon)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Berhasil menambah coupon",
	})
}

func UpdateCoupon(c *gin.Context) {

	couponId, errConv := strconv.Atoi(c.Param("id"))

	if errConv != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": errConv.Error(),
		})
		return
	}

	var couponFromJson models.Coupon

	err := utils.ValidateJsonFields(c, &couponFromJson)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	var defaultCoupon models.Coupon

	coupon, err := repositories.GetCouponById(database.DbConnection, couponId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	defaultCoupon.ID = coupon.ID
	defaultCoupon.Name = coupon.Name
	defaultCoupon.Discount = coupon.Discount

	utils.SetDefaultField(&couponFromJson, defaultCoupon)

	err = repositories.UpdateCoupon(database.DbConnection, couponFromJson)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"result": "Berhasil update coupon",
		})
	}
}

func DeleteCoupon(c *gin.Context) {

	couponId, errConv := strconv.Atoi(c.Param("id"))

	if errConv != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": errConv.Error(),
		})
		return
	}

	err := repositories.DeleteCoupon(database.DbConnection, couponId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Berhasil hapus coupon",
	})
}
