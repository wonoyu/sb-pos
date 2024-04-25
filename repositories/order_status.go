package repositories

import (
	"database/sql"
	"fmt"
	"sb-pos/models"
	"time"
)

func GetOrderStatus(db *sql.DB) (results []models.OrderStatus, err error) {
	sqlStmt := "SELECT * FROM order_status"

	rows, err := db.Query(sqlStmt)

	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		var orderStatus models.OrderStatus

		err = rows.Scan(&orderStatus.ID, &orderStatus.Name, &orderStatus.CreatedAt, &orderStatus.UpdatedAt)

		results = append(results, orderStatus)
	}

	return
}

func GetOrderStatusById(db *sql.DB, statusId int) (result models.OrderStatus, err error) {
	sqlStatement := "SELECT * FROM order_status WHERE id=$1"

	rows, errQuery := db.Query(sqlStatement, statusId)
	if errQuery != nil {
		err = errQuery
		return
	}

	if rows.Next() {
		var orderStatus = models.OrderStatus{}

		err = rows.Scan(&orderStatus.ID, &orderStatus.Name, &orderStatus.CreatedAt, &orderStatus.UpdatedAt)
		if err != nil {
			return
		}

		result = orderStatus
	} else {
		err = fmt.Errorf("there is no order status found with this identifier (%d)", statusId)
	}

	return
}

func InsertOrderStatus(db *sql.DB, orderStatus models.OrderStatus) (err error) {
	sql := "INSERT INTO order_status (name, created_at, updated_at)" +
		"VALUES ($1, $2, $3)"

	errs := db.QueryRow(sql, &orderStatus.Name, time.Now(), time.Now())

	return errs.Err()
}

func UpdateOrderStatus(db *sql.DB, orderStatus models.OrderStatus) (err error) {
	sqlStatement := "UPDATE order_status SET name=$1, updated_at=$2 " +
		"WHERE id=$3 Returning id"

	errs := db.QueryRow(sqlStatement, orderStatus.Name,
		time.Now(), orderStatus.ID).Scan(&orderStatus.ID)

	return errs
}

func DeleteOrderStatus(db *sql.DB, statusId int) (err error) {
	sql := "DELETE FROM order_status WHERE id=$1 Returning id"

	errs := db.QueryRow(sql, statusId).Scan(&statusId)

	return errs
}
