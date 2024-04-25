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

func GetAllUsers(c *gin.Context) {
	users, err := repositories.GetAllUsers(database.DbConnection)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"result": users,
		})
	}
}

func GetUserById(c *gin.Context) {
	userId, errConv := strconv.Atoi(c.Param("id"))

	if errConv != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": errConv.Error(),
		})
		return
	}

	user, err := repositories.GetUserById(database.DbConnection, userId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"result": user,
		})
	}
}

func UpdateUser(c *gin.Context) {
	userId, errConv := strconv.Atoi(c.Param("id"))

	if errConv != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": errConv.Error(),
		})
		return
	}

	var userFromJSON models.UpdateUser

	err := utils.ValidateJsonFields(c, &userFromJSON)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	var defaultUser models.UpdateUser

	user, err := repositories.GetUserById(database.DbConnection, userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	defaultUser.ID = user.ID
	defaultUser.Email = user.Email
	defaultUser.Username = user.Username

	utils.SetDefaultField(&userFromJSON, defaultUser)

	err = repositories.UpdateUser(database.DbConnection, userFromJSON)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"result": "Berhasil update user",
		})
	}
}

func UpdateUserRole(c *gin.Context) {

	userId, errConv := strconv.Atoi(c.Param("id"))

	if errConv != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": errConv.Error(),
		})
		return
	}

	var user models.UpdateUserRole

	err := utils.ValidateJsonFields(c, &user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	user.ID = userId

	err = repositories.UpdateUserRole(database.DbConnection, user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Berhasil update user role",
	})
}

func DeleteUser(c *gin.Context) {

	userId, errConv := strconv.Atoi(c.Param("id"))

	if errConv != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": errConv.Error(),
		})
		return
	}

	err := repositories.DeleteUser(database.DbConnection, userId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User berhasil dihapus",
	})
}
