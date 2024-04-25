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

func GetAllRoles(c *gin.Context) {

	roles, err := repositories.GetAllRoles(database.DbConnection)

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

func GetRoleById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	role, err := repositories.GetRoleById(database.DbConnection, id)

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

func InsertRole(c *gin.Context) {

	var role models.Role

	err := utils.ValidateJsonFields(c, &role)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	err = repositories.InsertRole(database.DbConnection, role)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Berhasil menambah role",
	})
}

func UpdateRole(c *gin.Context) {

	roleId, errConv := strconv.Atoi(c.Param("id"))

	if errConv != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": errConv.Error(),
		})
		return
	}

	var role models.Role

	err := utils.ValidateJsonFields(c, &role)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	role.ID = roleId

	err = repositories.UpdateRole(database.DbConnection, role)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Berhasil update role",
	})
}

func DeleteRole(c *gin.Context) {
	roleId, errConv := strconv.Atoi(c.Param("id"))

	if errConv != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": errConv.Error(),
		})
		return
	}

	err := repositories.DeleteRole(database.DbConnection, roleId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Berhasil hapus role",
	})
}
