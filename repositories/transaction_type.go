package repositories

import (
	"database/sql"
	"fmt"
	"sb-pos/models"
	"time"
)

func GetTransactionType(db *sql.DB) (results []models.TransactionType, err error) {
	sqlStmt := "SELECT * FROM transaction_type"

	rows, err := db.Query(sqlStmt)

	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		var transactionType models.TransactionType

		err = rows.Scan(&transactionType.ID, &transactionType.Name, &transactionType.CreatedAt, &transactionType.UpdatedAt)

		results = append(results, transactionType)
	}

	return
}

func GetTransactionTypeById(db *sql.DB, typeId int) (result models.TransactionType, err error) {
	sqlStatement := "SELECT * FROM transaction_type WHERE id=$1"

	rows, errQuery := db.Query(sqlStatement, typeId)
	if errQuery != nil {
		err = errQuery
		return
	}

	if rows.Next() {
		var transactionType = models.TransactionType{}

		err = rows.Scan(&transactionType.ID, &transactionType.Name, &transactionType.CreatedAt, &transactionType.UpdatedAt)
		if err != nil {
			return
		}

		result = transactionType
	} else {
		err = fmt.Errorf("there is no transaction type found with this identifier (%d)", typeId)
	}

	return
}

func InsertTransactionType(db *sql.DB, transactionType models.TransactionType) (err error) {
	sql := "INSERT INTO transaction_type (name, created_at, updated_at)" +
		"VALUES ($1, $2, $3)"

	errs := db.QueryRow(sql, &transactionType.Name, time.Now(), time.Now())

	return errs.Err()
}

func UpdateTransactionType(db *sql.DB, transactionType models.TransactionType) (err error) {
	sqlStatement := "UPDATE transaction_type SET name=$1, updated_at=$2 " +
		"WHERE id=$3 Returning id"

	errs := db.QueryRow(sqlStatement, transactionType.Name,
		time.Now(), transactionType.ID).Scan(&transactionType.ID)

	return errs
}

func DeleteTransactionType(db *sql.DB, typeId int) (err error) {
	sql := "DELETE FROM transaction_type WHERE id=$1 Returning id"

	errs := db.QueryRow(sql, typeId).Scan(&typeId)

	return errs
}
