package utils

import (
	"errors"
	"os"
	"sb-pos/constants"
	"sb-pos/models"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type Claims struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	UserId   string `json:"user_id"`
	RoleName string `json:"role_name"`
	jwt.StandardClaims
}

func GenerateToken(user models.UserWithPassword) (token string, err error) {
	c := Claims{
		user.Username,
		user.Email,
		strconv.Itoa(user.ID),
		user.RoleName,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(constants.TokenExpireDuration).Unix(),
			Issuer:    "sb-pos",
		},
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	return t.SignedString([]byte(os.Getenv("SECRET_KEY")))
}

func parseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (i interface{}, err error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("token is invalid")
}

func GetToken(c *gin.Context) (claims *Claims, err error) {
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		return nil, errors.New("unauthorized (no token were found)")
	}

	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		return nil, errors.New("invalid authorization format")
	}

	claims, err = parseToken(parts[1])
	if err != nil {
		return nil, errors.New("invalid token")
	}

	return
}
