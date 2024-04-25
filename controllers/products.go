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

func GetProducts(c *gin.Context) {
	products, err := repositories.GetProducts(database.DbConnection)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": products,
	})
}

func GetProductById(c *gin.Context) {
	productId, errConv := strconv.Atoi(c.Param("id"))

	if errConv != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": errConv.Error(),
		})
		return
	}

	product, err := repositories.GetProductById(database.DbConnection, productId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": product,
	})
}

func InsertProduct(c *gin.Context) {

	var product models.CreateProduct

	err := utils.ValidateJsonFields(c, &product)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	err = repositories.InsertProduct(database.DbConnection, product)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Berhasil menambah product",
	})
}

func UpdateProduct(c *gin.Context) {
	productId, errConv := strconv.Atoi(c.Param("id"))

	if errConv != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": errConv.Error(),
		})
		return
	}

	var productFromJSON models.Product

	err := utils.ValidateJsonFields(c, &productFromJSON)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	var defaultProduct models.Product

	product, err := repositories.GetProductById(database.DbConnection, productId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	defaultProduct.ID = product.ID
	defaultProduct.Name = product.Name
	defaultProduct.Price = product.Price
	defaultProduct.StockQuantity = product.StockQuantity
	defaultProduct.CategoryId = product.CategoryId

	utils.SetDefaultField(&productFromJSON, defaultProduct)

	err = repositories.UpdateProduct(database.DbConnection, productFromJSON)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"result": "Berhasil update product",
		})
	}
}

func DeleteProduct(c *gin.Context) {

	productId, errConv := strconv.Atoi(c.Param("id"))

	if errConv != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": errConv.Error(),
		})
		return
	}

	err := repositories.DeleteProduct(database.DbConnection, productId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Berhasil hapus product",
	})
}
