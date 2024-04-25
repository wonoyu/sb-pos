package controllers

import (
	"net/http"
	"sb-pos/database"
	"sb-pos/repositories"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetSalesTransactions(c *gin.Context) {
	transactions, err := repositories.GetSalesTransactions(database.DbConnection)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"results": transactions,
	})
}

func GetSalesTransactionById(c *gin.Context) {
	transactionId, errConv := strconv.Atoi(c.Param("id"))

	if errConv != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": errConv.Error(),
		})
		return
	}

	transaction, err := repositories.GetSalesTransactionsById(database.DbConnection, transactionId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": transaction,
	})
}
