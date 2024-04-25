package controllers

import (
	"net/http"
	"sb-pos/database"
	"sb-pos/models"
	"sb-pos/repositories"
	"sb-pos/utils"

	"github.com/gin-gonic/gin"
)

func GetProfile(c *gin.Context) {
	userId, err := utils.GetUserId(c)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	customer, err := repositories.GetCustomerByUserId(database.DbConnection, userId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": customer,
	})
}

func TopUp(c *gin.Context) {
	userId, err := utils.GetUserId(c)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	customer, err := repositories.GetCustomerByUserId(database.DbConnection, userId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	var order models.TopupBalance

	err = utils.ValidateJsonFields(c, &order)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	order.CustomerId = customer.ID

	err = repositories.UpdateBalance(database.DbConnection, order)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "topup berhasil",
	})
}
