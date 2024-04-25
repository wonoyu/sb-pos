package repositories

import (
	"database/sql"
	"fmt"
	"sb-pos/models"
	"time"
)

func RegisterCustomer(db *sql.DB, userId int) (err error) {
	sql := "INSERT INTO customers (balance, user_id, created_at, updated_at)" +
		"VALUES ($1, $2, $3, $4)"

	errs := db.QueryRow(sql, 0, userId, time.Now(), time.Now())

	return errs.Err()
}

func GetCustomerByUserId(db *sql.DB, userId int) (result models.Customer, err error) {
	sqlStatement := "SELECT c.id, c.balance, u.username, u.email " +
		"FROM customers c JOIN users u ON c.user_id=u.id WHERE u.id=$1"

	rows, errQuery := db.Query(sqlStatement, userId)
	if errQuery != nil {
		err = errQuery
		return
	}

	if rows.Next() {
		var customer = models.Customer{}

		err = rows.Scan(&customer.ID, &customer.Balance, &customer.Username, &customer.Email)
		if err != nil {
			return
		}

		result = customer
	} else {
		err = fmt.Errorf("there is no customer found with this identifier (%d)", userId)
	}

	return
}

func UpdateBalance(db *sql.DB, customer models.TopupBalance) (err error) {
	sql := "UPDATE customers SET balance=balance + $1, updated_at=$2" +
		"WHERE id=$3 Returning id"

	balance := customer.Balance

	if balance < 0 {
		balance = 0
	}

	errs := db.QueryRow(sql, balance, time.Now(), customer.CustomerId).Scan(&customer.CustomerId)

	return errs
}

func GetCustomerOrder(db *sql.DB, customerId int) (results []models.CustomerOrder, err error) {
	sqlStmt := "SELECT co.id, co.customer_id, o.name FROM customer_orders co JOIN orders o ON co.order_id=o.id " +
		"WHERE co.customer_id=$2 ORDER BY o.name ASC"

	rows, err := db.Query(sqlStmt, customerId)
	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		var order = models.CustomerOrder{}

		err = rows.Scan(&order.ID, &order.CustomerId, &order.OrderName)
		if err != nil {
			return
		}

		products, err := GetOrderProductsById(db, order.ID)

		if err != nil {
			break
		}

		var totalPrice int64 = 0
		for _, product := range products {
			totalPrice += (product.ProductPrice * int64(product.ProductQuantity))
		}

		order.Products = products

		results = append(results, order)
	}

	return
}

func GetCustomerOrderByOrderId(db *sql.DB, orderId int) (result models.CustomerOrder, err error) {
	products, err := GetOrderProductsById(db, orderId)

	if err != nil {
		return
	}

	var totalPrice int64 = 0
	for _, product := range products {
		totalPrice += (product.ProductPrice * int64(product.ProductQuantity))
	}

	sqlStmt := "SELECT co.id, co.customer_id, o.name FROM customer_orders co JOIN orders o ON co.order_id=o.id " +
		"WHERE co.order_id=$1 ORDER BY o.name ASC"

	rows, err := db.Query(sqlStmt, orderId)
	if err != nil {
		return
	}

	defer rows.Close()

	if rows.Next() {
		var order = models.CustomerOrder{}

		err = rows.Scan(&order.ID, &order.CustomerId, &order.OrderName)
		if err != nil {
			return
		}

		order.TotalPrice = totalPrice

		result = order
	} else {
		err = fmt.Errorf("there is no product category found with this identifier (%d)", orderId)
	}

	return
}

func CreateCustomerOrder(db *sql.DB, orderId int, customerId int) (err error) {
	sqlStmt := "INSERT INTO customer_orders (order_id, customer_id, created_at, updated_at) " +
		"VALUES ($1, $2, $3, $4)"

	errs := db.QueryRow(sqlStmt, orderId, customerId, time.Now(), time.Now())

	return errs.Err()
}
