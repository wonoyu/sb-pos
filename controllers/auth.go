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

func registerUser(c *gin.Context) {
	var user models.RegisterUser

	err := utils.ValidateJsonFields(c, &user)

	if id, exists := c.Get("role_id"); exists {
		roleId, err := strconv.Atoi(id.(string))

		if roleId == 0 {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "role_id tidak ditemukan",
			})
			return
		}

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		user.RoleId = roleId
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	if !utils.ValidateEmail(database.DbConnection, user.Email) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Email tidak dapat digunakan",
		})
		return
	}

	encrypted, err := utils.EncryptPassword(user.Password)
	user.Password = encrypted

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	err = repositories.RegisterUser(database.DbConnection, user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	// register customer
	if user.RoleId == 2 {
		result, err := repositories.GetUserByEmail(database.DbConnection, user.Email)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		err = repositories.RegisterCustomer(database.DbConnection, result.ID)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Registrasi Berhasil",
	})
}

func RegisterCashier(c *gin.Context) {
	c.Set("role_id", "3")

	registerUser(c)
}

func RegisterCustomer(c *gin.Context) {
	c.Set("role_id", "2")

	registerUser(c)
}

func RegisterAdmin(c *gin.Context) {
	c.Set("role_id", "1")

	registerUser(c)
}

func Login(c *gin.Context) {
	var login models.Login

	err := c.ShouldBindJSON(&login)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	user, err := repositories.GetUserByEmail(database.DbConnection, login.Email)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
		})
		return
	}

	validatedPw, err := utils.ValidatePassword(login.Password, user.Password)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
		})
		return
	}

	user.Password = validatedPw

	tokenString, err := utils.GenerateToken(user)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}
