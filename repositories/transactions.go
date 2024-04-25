package repositories

import (
	"database/sql"
	"fmt"
	"sb-pos/models"
)

func GetSalesTransactions(db *sql.DB) (results []models.SalesTransaction, err error) {
	sqlStmt := "SELECT t.id, t.order_id, o.name, t.customer_id, u.username, t.created_at, t.updated_at " +
		"FROM transactions t JOIN customers c ON t.customer_id=c.id " +
		"JOIN users u ON c.user_id = u.id " +
		"JOIN orders o ON t.order_id = o.id ORDER BY t.created_at DESC"

	rows, err := db.Query(sqlStmt)

	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		var transaction models.SalesTransaction

		err = rows.Scan(&transaction.ID, &transaction.OrderId, &transaction.OrderName, &transaction.CustomerId,
			&transaction.CustomerName, &transaction.CreatedAt, &transaction.UpdatedAt)

		if err != nil {
			return
		}

		products, err := GetOrderProductsById(db, transaction.OrderId)

		if err != nil {
			break
		}

		transaction.Products = products

		results = append(results, transaction)
	}

	return
}

func GetSalesTransactionsById(db *sql.DB, transactionId int) (result models.SalesTransaction, err error) {
	sqlStmt := "SELECT t.id, t.order_id, o.name, t.customer_id, u.username, t.created_at, t.updated_at " +
		"FROM transactions t JOIN customers c ON t.customer_id=c.id " +
		"JOIN users u ON c.user_id = u.id " +
		"JOIN orders o ON t.order_id = o.id WHERE t.id=$1"

	rows, err := db.Query(sqlStmt, transactionId)

	if err != nil {
		return
	}

	defer rows.Close()

	if rows.Next() {
		var transaction models.SalesTransaction

		err = rows.Scan(&transaction.ID, &transaction.OrderId, &transaction.OrderName, &transaction.CustomerId,
			&transaction.CustomerName, &transaction.CreatedAt, &transaction.UpdatedAt)

		if err != nil {
			return
		}

		products, err := GetOrderProductsById(db, transaction.OrderId)

		if err != nil {
			return result, err
		}

		transaction.Products = products

		result = transaction
	} else {
		err = fmt.Errorf("there is no transaction found with this identifier (%d)", transactionId)
	}

	return
}
