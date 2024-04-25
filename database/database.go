package database

import (
	"database/sql"
	"embed"
	"fmt"
	"os"
	"sb-pos/models"
	"sb-pos/repositories"
	"sb-pos/utils"

	migrate "github.com/rubenv/sql-migrate"
)

var (
	DbConnection *sql.DB
)

//go:embed migrations/*.sql
var dbMigrations embed.FS

func DbMigrate(dbParam *sql.DB) {

	migrations := &migrate.EmbedFileSystemMigrationSource{
		FileSystem: dbMigrations,
		Root:       "migrations",
	}

	n, errs := migrate.Exec(dbParam, "postgres", migrations, migrate.Up)
	if errs != nil {
		panic(errs)
	}

	DbConnection = dbParam

	errs = createAdmin()

	if errs != nil {
		panic(errs)
	}

	errs = createCustomer()

	if errs != nil {
		panic(errs)
	}

	errs = createCashier()

	if errs != nil {
		panic(errs)
	}

	fmt.Println("Applied", n, "migrations!")
}

func createAdmin() (err error) {
	pwd, err := utils.EncryptPassword(os.Getenv("ADM_PWD"))

	if err != nil {
		return
	}

	admin := models.RegisterUser{
		Username: "admin",
		Email:    "admin@mail.com",
		Password: pwd,
		RoleId:   1,
	}

	if !utils.ValidateEmail(DbConnection, admin.Email) {
		return
	}

	err = repositories.RegisterUser(DbConnection, admin)

	return err
}

func createCustomer() (err error) {
	pwd, err := utils.EncryptPassword(os.Getenv("CUSTOMER_PWD"))

	if err != nil {
		return
	}

	customer := models.RegisterUser{
		Username: "customer",
		Email:    "customer@mail.com",
		Password: pwd,
		RoleId:   2,
	}

	if !utils.ValidateEmail(DbConnection, customer.Email) {
		return
	}

	err = repositories.RegisterUser(DbConnection, customer)

	if err != nil {
		return
	}

	err = repositories.RegisterCustomer(DbConnection, 2)

	return err
}

func createCashier() (err error) {
	pwd, err := utils.EncryptPassword(os.Getenv("CASHIER_PWD"))

	if err != nil {
		return
	}

	cashier := models.RegisterUser{
		Username: "cashier",
		Email:    "cashier@mail.com",
		Password: pwd,
		RoleId:   3,
	}

	if !utils.ValidateEmail(DbConnection, cashier.Email) {
		return
	}

	err = repositories.RegisterUser(DbConnection, cashier)

	return err
}
