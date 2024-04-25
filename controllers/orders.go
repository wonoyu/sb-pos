package controllers

import (
	"fmt"
	"net/http"
	"sb-pos/database"
	"sb-pos/models"
	"sb-pos/repositories"
	"sb-pos/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetOrders(c *gin.Context) {
	orders, err := repositories.GetOrders(database.DbConnection)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": orders,
	})
}

func CreateOrder(c *gin.Context) {
	userId, err := utils.GetUserId(c)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	var order models.CreateOrder

	err = utils.ValidateJsonFields(c, &order)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	for _, product := range order.Products {
		p, err := repositories.GetProductById(database.DbConnection, product.ProductId)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		if p.StockQuantity < product.Quantity {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": fmt.Sprintf("Stock produk (%s) tidak cukup, harap sesuaikan quantity produk "+
					"menjadi kurang atau sama dengan %d", p.Name, p.StockQuantity),
			})
			return
		}
	}

	customer, err := repositories.GetCustomerByUserId(database.DbConnection, userId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	orderId, err := repositories.CreateOrder(database.DbConnection, order)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	err = repositories.CreateOrderProducts(database.DbConnection, orderId, order.Products)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	err = repositories.CreateCustomerOrder(database.DbConnection, orderId, customer.ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Berhasil membuat order",
	})
}

func PayOrder(c *gin.Context) {
	userId, err := utils.GetUserId(c)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	orderId, err := strconv.Atoi(c.Param("id"))

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

	totalPrice, err := repositories.GetOrderProductsTotalPrice(database.DbConnection, orderId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	if (customer.Balance - totalPrice) < 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "saldo tidak cukup, silahkan topup",
		})
		return
	}

	err = repositories.PayOrder(database.DB, orderId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	err = repositories.UpdateBalance(database.DbConnection, models.TopupBalance{
		CustomerId: customer.ID,
		Balance:    customer.Balance - totalPrice,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Pembayaran telah diterima, silahkan cek di kasir",
	})
}

func CompleteOrder(c *gin.Context) {
	orderId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	orderById, err := repositories.GetOrderById(database.DbConnection, orderId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	if orderById.StatusName == "unpaid" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "order ini belum dibayar",
		})
		return
	}

	var order models.CompleteOrder

	customerOrder, err := repositories.GetCustomerOrderByOrderId(database.DbConnection, orderId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "get order " + err.Error(),
		})
		return
	}

	order.OrderId = orderId
	order.CustomerId = customerOrder.CustomerId

	err = repositories.CompleteOrder(database.DbConnection, order)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "complete " + err.Error(),
		})
		return
	}

	for _, prdct := range orderById.Products {
		err = repositories.UpdateProductQuantity(database.DbConnection, prdct.ProductId, prdct.ProductQuantity)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "uqty " + err.Error(),
			})
			return
		}
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Order berhasil diselesaikan",
	})
}
