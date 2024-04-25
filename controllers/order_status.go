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

func GetOrderStatus(c *gin.Context) {
	orderStatus, err := repositories.GetOrderStatus(database.DbConnection)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": orderStatus,
	})
}

func GetOrderStatusById(c *gin.Context) {
	statusId, errConv := strconv.Atoi(c.Param("id"))

	if errConv != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": errConv.Error(),
		})
		return
	}

	orderStatus, err := repositories.GetOrderStatusById(database.DbConnection, statusId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": orderStatus,
	})
}

func InsertOrderStatus(c *gin.Context) {

	var orderStatus models.OrderStatus

	err := utils.ValidateJsonFields(c, &orderStatus)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	err = repositories.InsertOrderStatus(database.DbConnection, orderStatus)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Berhasil menambah order status",
	})
}

func UpdateOrderStatus(c *gin.Context) {

	statusId, errConv := strconv.Atoi(c.Param("id"))

	if errConv != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": errConv.Error(),
		})
		return
	}

	var orderStatus models.OrderStatus

	err := utils.ValidateJsonFields(c, &orderStatus)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	orderStatus.ID = statusId

	err = repositories.UpdateOrderStatus(database.DbConnection, orderStatus)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Berhasil update order status",
	})
}

func DeleteOrderStatus(c *gin.Context) {

	statusId, errConv := strconv.Atoi(c.Param("id"))

	if errConv != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": errConv.Error(),
		})
		return
	}

	err := repositories.DeleteOrderStatus(database.DbConnection, statusId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Berhasil hapus order status",
	})
}
