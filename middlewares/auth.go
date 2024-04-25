package middlewares

import (
	"net/http"
	"sb-pos/utils"

	"github.com/gin-gonic/gin"
)

func JWTAuth() func(c *gin.Context) {
	return func(c *gin.Context) {
		claims, err := utils.GetToken(c)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": err.Error(),
			})
			c.Abort()
			return
		}

		c.Set("username", claims.Username)
		c.Set("email", claims.Email)
		c.Set("user_id", claims.UserId)
		c.Set("role_name", claims.RoleName)

		c.Next()
	}
}

func JWTAuthAdmin() func(c *gin.Context) {
	return func(c *gin.Context) {
		claims, err := utils.GetToken(c)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": err.Error(),
			})
			c.Abort()
			return
		}

		c.Set("username", claims.Username)
		c.Set("email", claims.Email)
		c.Set("user_id", claims.UserId)
		c.Set("role_name", claims.RoleName)

		if claims.RoleName != "admin" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "only admin can access this resource",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

func JWATAuthCustomer() func(c *gin.Context) {
	return func(c *gin.Context) {
		claims, err := utils.GetToken(c)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": err.Error(),
			})
			c.Abort()
			return
		}

		c.Set("username", claims.Username)
		c.Set("email", claims.Email)
		c.Set("user_id", claims.UserId)
		c.Set("role_name", claims.RoleName)

		if claims.RoleName != "customer" && claims.RoleName != "admin" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "only customer and admin can access this resource",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

func JWATAuthCashier() func(c *gin.Context) {
	return func(c *gin.Context) {
		claims, err := utils.GetToken(c)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": err.Error(),
			})
			c.Abort()
			return
		}

		c.Set("username", claims.Username)
		c.Set("email", claims.Email)
		c.Set("user_id", claims.UserId)
		c.Set("role_name", claims.RoleName)

		if claims.RoleName != "cashier" && claims.RoleName != "admin" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "only cashier and admin can access this resource",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
