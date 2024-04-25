package repositories

import (
	"database/sql"
	"fmt"
	"sb-pos/models"
	"time"
)

func GetProductCategories(db *sql.DB) (results []models.ProductCategory, err error) {
	sqlStmt := "SELECT * FROM product_category"

	rows, err := db.Query(sqlStmt)

	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		var category models.ProductCategory

		err = rows.Scan(&category.ID, &category.Name, &category.CreatedAt, &category.UpdatedAt)

		results = append(results, category)
	}

	return
}

func GetProductCategoryById(db *sql.DB, categoryId int) (result models.ProductCategory, err error) {
	sqlStatement := "SELECT * FROM product_category WHERE id=$1"

	rows, errQuery := db.Query(sqlStatement, categoryId)
	if errQuery != nil {
		err = errQuery
		return
	}

	if rows.Next() {
		var category = models.ProductCategory{}

		err = rows.Scan(&category.ID, &category.Name, &category.CreatedAt, &category.UpdatedAt)
		if err != nil {
			return
		}

		result = category
	} else {
		err = fmt.Errorf("there is no product category found with this identifier (%d)", categoryId)
	}

	return
}

func InsertProductCategory(db *sql.DB, category models.ProductCategory) (err error) {
	sql := "INSERT INTO product_category (name, created_at, updated_at)" +
		"VALUES ($1, $2, $3)"

	errs := db.QueryRow(sql, &category.Name, time.Now(), time.Now())

	return errs.Err()
}

func UpdateProductCategory(db *sql.DB, category models.ProductCategory) (err error) {
	sqlStatement := "UPDATE product_category SET name=$1, updated_at=$2 " +
		"WHERE id=$3 Returning id"

	errs := db.QueryRow(sqlStatement, category.Name,
		time.Now(), category.ID).Scan(&category.ID)

	return errs
}

func DeleteProductCategory(db *sql.DB, categoryId int) (err error) {
	sql := "DELETE FROM product_category WHERE id=$1 Returning id"

	errs := db.QueryRow(sql, categoryId).Scan(&categoryId)

	return errs
}
