package utils

import (
	"database/sql"
	"encoding/hex"
	"errors"
	"fmt"
	"net/mail"
	"reflect"
	"sb-pos/repositories"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

func ValidateJsonFields(c *gin.Context, obj interface{}) (err error) {
	validate := validator.New()

	if err = c.ShouldBindJSON(obj); err != nil {
		return
	}

	if err = validate.Struct(obj); err != nil {
		return
	}

	return
}

func EncryptPassword(password string) (encrypted string, err error) {
	hash, errs := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return fmt.Sprintf("%x", hash), errs
}

func ValidateEmail(db *sql.DB, email string) bool {
	_, err := mail.ParseAddress(email)

	if err != nil {
		return false
	}

	_, err = repositories.GetUserByEmail(db, email)

	return err != nil
}

func ValidatePassword(password string, storedHexHash string) (decrypted string, err error) {
	storedHash, err := hex.DecodeString(storedHexHash)

	if err != nil {
		return
	}

	err = bcrypt.CompareHashAndPassword(storedHash, []byte(password))

	return password, err
}

func SetDefaultField(target, data interface{}) {
	targetValue := reflect.ValueOf(target).Elem()
	dataValue := reflect.ValueOf(data)

	for i := 0; i < targetValue.NumField(); i++ {
		targetField := targetValue.Field(i)
		dataField := dataValue.Field(i)

		if targetField.Interface() == reflect.Zero(targetField.Type()).Interface() {
			targetField.Set(dataField)
		}
	}
}

func GetUserId(c *gin.Context) (id int, err error) {
	_, exists := c.Get("user_id")

	if !exists {
		err = errors.New("user id cannot be found")
		return
	}

	userId, err := strconv.Atoi(c.GetString("user_id"))

	if err != nil {
		return
	}

	id = userId

	return
}

func LoadEnv() {
	err := godotenv.Load("config/.env")

	if err != nil {
		fmt.Println("failed load file environment")
		panic(err)
	} else {
		fmt.Println("success load file environment")
	}
}
