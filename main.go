package main

import (
	"sb-pos/database"
	"sb-pos/router"
	"sb-pos/utils"
)

// TODO: Bussiness logic transaksi (orders, customers, log_transaksi)
// TODO: customers -> create order, pay order, refund order, can check orders
// TODO: transactions -> used to calculate refunds and sales, tipe order harus paid
// TODO: payment -> jika balance customer tidak cukup (qty * price) maka lempar error
// TODO: orders -> generate order name format order_P1_date

func main() {
	utils.LoadEnv()
	database.ConnectDb()
	router := router.InitRouter()

	router.Run()
}
