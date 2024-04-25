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

func GetTransactionType(c *gin.Context) {
	orderStatus, err := repositories.GetTransactionType(database.DbConnection)

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

func GetTransactionTypeById(c *gin.Context) {
	typeId, errConv := strconv.Atoi(c.Param("id"))

	if errConv != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": errConv.Error(),
		})
		return
	}

	transactionType, err := repositories.GetTransactionTypeById(database.DbConnection, typeId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": transactionType,
	})
}

func InsertTransactionType(c *gin.Context) {

	var transactionType models.TransactionType

	err := utils.ValidateJsonFields(c, &transactionType)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	err = repositories.InsertTransactionType(database.DbConnection, transactionType)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Berhasil menambah transaction type",
	})
}

func UpdateTransactionType(c *gin.Context) {

	typeId, errConv := strconv.Atoi(c.Param("id"))

	if errConv != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": errConv.Error(),
		})
		return
	}

	var transactionType models.TransactionType

	err := utils.ValidateJsonFields(c, &transactionType)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	transactionType.ID = typeId

	err = repositories.UpdateTransactionType(database.DbConnection, transactionType)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Berhasil update transaction type",
	})
}

func DeleteTransactionType(c *gin.Context) {

	typeId, errConv := strconv.Atoi(c.Param("id"))

	if errConv != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": errConv.Error(),
		})
		return
	}

	err := repositories.DeleteTransactionType(database.DbConnection, typeId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Berhasil hapus transaction type",
	})
}
