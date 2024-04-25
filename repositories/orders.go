package repositories

import (
	"database/sql"
	"fmt"
	"sb-pos/models"
	"time"
)

func GetOrders(db *sql.DB) (results []models.Order, err error) {
	sqlStmt := "SELECT o.id, o.name, o.status_id, s.name as status, o.created_at, o.updated_at " +
		"FROM orders o JOIN order_status s ON o.status_id = s.id " +
		"ORDER BY o.created_at DESC"

	rows, err := db.Query(sqlStmt)
	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		var order = models.Order{}

		err = rows.Scan(&order.ID, &order.Name, &order.StatusId, &order.StatusName,
			&order.CreatedAt, &order.UpdatedAt)

		if err != nil {
			return
		}

		products, err := GetOrderProductsById(db, order.ID)

		if err != nil {
			break
		}

		order.Products = products

		results = append(results, order)
	}

	return
}

func GetOrderById(db *sql.DB, orderId int) (result models.Order, err error) {
	sqlStatement := "SELECT o.id, o.name, o.status_id ,s.name as status, o.created_at, o.updated_at " +
		"FROM orders o JOIN order_status s ON o.status_id = s.id " +
		"WHERE o.id = $1"

	rows, errQuery := db.Query(sqlStatement, orderId)
	if errQuery != nil {
		err = errQuery
		return
	}

	if rows.Next() {
		var order = models.Order{}

		err = rows.Scan(&order.ID, &order.Name, &order.StatusId, &order.StatusName,
			&order.CreatedAt, &order.UpdatedAt)

		if err != nil {
			return
		}

		products, err := GetOrderProductsById(db, order.ID)

		if err != nil {
			return result, err
		}

		order.Products = products

		result = order
	} else {
		err = fmt.Errorf("there is no order found with this identifier (%d)", orderId)
	}

	return
}

func CreateOrder(db *sql.DB, order models.CreateOrder) (orderId int, err error) {
	sql := "INSERT INTO orders (name, status_id, created_at, updated_at)" +
		"VALUES ($1, $2, $3, $4) RETURNING id"

	var id int
	year, month, day := time.Now().Date()
	err = db.QueryRow(sql, fmt.Sprintf("order_%d-%d-%d", year, month, day), 2, time.Now(), time.Now()).Scan(&id)

	orderId = id

	if orderId == id {
		return orderId, nil
	}

	fmt.Printf("order %s", err.Error())

	return
}

func CreateOrderProducts(db *sql.DB, orderId int, products []models.OrderProducts) (err error) {
	for _, product := range products {
		sqlStmt := "INSERT INTO order_products (order_id, product_id, product_quantity, created_at, updated_at) " +
			"VALUES ($1, $2, $3, $4, $5)"

		errs := db.QueryRow(sqlStmt, orderId, product.ProductId, product.Quantity, time.Now(), time.Now())

		err = errs.Err()
	}

	return
}

func GetOrderProductsById(db *sql.DB, orderId int) (results []models.OrderProductComplete, err error) {
	sqlStmt := "SELECT op.id, op.product_id, p.name, op.product_quantity, p.price, op.created_at, op.updated_at " +
		"FROM order_products op JOIN products p ON op.product_id = p.id " +
		"WHERE op.order_id = $1"

	rows, err := db.Query(sqlStmt, orderId)

	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		var product models.OrderProductComplete

		err = rows.Scan(&product.ID, &product.ProductId, &product.ProductName, &product.ProductQuantity,
			&product.ProductPrice, &product.CreatedAt, &product.UpdatedAt)

		if err != nil {
			return
		}

		results = append(results, product)
	}

	return
}

func GetOrderProductsTotalPrice(db *sql.DB, orderId int) (totalPrice int64, err error) {
	sqlStmt := "SELECT SUM(op.product_quantity * p.price) as total_price " +
		"FROM order_products op JOIN products p ON op.product_id=p.id " +
		"WHERE op.order_id = $1"

	rows, err := db.Query(sqlStmt, orderId)

	if err != nil {
		return
	}

	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(&totalPrice)

		if err != nil {
			return
		}
	} else {
		err = fmt.Errorf("there is no order found with this identifier (%d)", orderId)
	}

	return
}

func PayOrder(db *sql.DB, orderId int) (err error) {
	sqlStatement := "UPDATE orders SET status_id=$1, updated_at=$2 " +
		"WHERE id=$3 Returning id"

	errs := db.QueryRow(sqlStatement, 1,
		time.Now(), orderId).Scan(&orderId)

	return errs
}

func CompleteOrder(db *sql.DB, order models.CompleteOrder) (err error) {
	sqlStatement := "INSERT INTO transactions (order_id, customer_id, ts_type_id, created_at, updated_at) " +
		"VALUES ($1, $2, $3, $4, $5)"

	errs := db.QueryRow(sqlStatement, order.OrderId, order.CustomerId, 1, time.Now(), time.Now())

	return errs.Err()
}

func DeleteOrder(db *sql.DB, orderId int) (err error) {
	sql := "DELETE FROM orders WHERE id=$1 Returning id"

	errs := db.QueryRow(sql, orderId).Scan(&orderId)

	return errs
}
