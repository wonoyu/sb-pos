package repositories

import (
	"database/sql"
	"fmt"
	"sb-pos/models"
	"time"
)

func GetProducts(db *sql.DB) (results []models.Product, err error) {
	sqlStmt := "SELECT p.*, c.name FROM products p JOIN product_category c " +
		"ON p.category_id=c.id ORDER BY p.created_at DESC"

	rows, err := db.Query(sqlStmt)

	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		var product models.Product

		err = rows.Scan(&product.ID, &product.Name, &product.Price, &product.StockQuantity,
			&product.CategoryId, &product.CreatedAt, &product.UpdatedAt, &product.CategoryName)

		if err != nil {
			return
		}

		results = append(results, product)
	}

	return
}

func GetProductById(db *sql.DB, productId int) (result models.Product, err error) {
	sqlStatement := "SELECT p.*, c.Name as category_name FROM products p JOIN product_category c " +
		"ON p.category_id=c.id WHERE p.id=$1"

	rows, errQuery := db.Query(sqlStatement, productId)
	if errQuery != nil {
		err = errQuery
		return
	}

	if rows.Next() {
		var product = models.Product{}

		err = rows.Scan(&product.ID, &product.Name, &product.Price, &product.StockQuantity,
			&product.CategoryId, &product.CreatedAt, &product.UpdatedAt, &product.CategoryName)
		if err != nil {
			return
		}

		result = product
	} else {
		err = fmt.Errorf("there is no product found with this identifier (%d)", productId)
	}

	return
}

func InsertProduct(db *sql.DB, product models.CreateProduct) (err error) {
	sql := "INSERT INTO products (name, price, stock_quantity, category_id, created_at, updated_at)" +
		"VALUES ($1, $2, $3, $4, $5, $6)"

	errs := db.QueryRow(sql, &product.Name, &product.Price, &product.StockQuantity, &product.CategoryId,
		time.Now(), time.Now())

	return errs.Err()
}

func UpdateProduct(db *sql.DB, product models.Product) (err error) {
	sqlStatement := "UPDATE products SET name=$1, price=$2, stock_quantity=$3, category_id=$4, " +
		"updated_at=$5 WHERE id=$6 Returning products.id"

	errs := db.QueryRow(sqlStatement, product.Name, product.Price, product.StockQuantity,
		product.CategoryId, time.Now(), product.ID).Scan(&product.ID)

	return errs
}

func UpdateProductQuantity(db *sql.DB, productId int, productQuantity int) (err error) {
	sqlStatement := "SELECT stock_quantity FROM products WHERE id=$1"

	rows, err := db.Query(sqlStatement, productId)

	if err != nil {
		return
	}

	defer rows.Close()

	var currentStock int
	if rows.Next() {
		err = rows.Scan(&currentStock)

		if err != nil {
			return
		}
	} else {
		err = fmt.Errorf("there is no product found with this identifier (%d)", productId)
	}

	sqlStmt := "UPDATE products SET stock_quantity=$1, updated_at=$2 WHERE id=$3"

	errs := db.QueryRow(sqlStmt, currentStock-productQuantity, time.Now(), productId)

	return errs.Err()
}

func DeleteProduct(db *sql.DB, productId int) (err error) {
	sql := "DELETE FROM products WHERE id=$1 Returning id"

	errs := db.QueryRow(sql, productId).Scan(&productId)

	return errs
}
