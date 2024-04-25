package repositories

import (
	"database/sql"
	"fmt"
	"sb-pos/models"
	"time"
)

func GetCoupons(db *sql.DB) (results []models.Coupon, err error) {
	sqlStmt := "SELECT * FROM coupons ORDER BY created_at DESC"

	rows, err := db.Query(sqlStmt)

	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		var coupon models.Coupon
		err = rows.Scan(&coupon.ID, &coupon.Name, &coupon.CreatedAt, &coupon.UpdatedAt)

		if err != nil {
			return
		}

		results = append(results, coupon)
	}

	return
}

func GetCouponById(db *sql.DB, couponId int) (result models.Coupon, err error) {
	sqlStmt := "SELECT * FROM coupons WHERE id=$1 ORDER BY created_at DESC"

	rows, err := db.Query(sqlStmt, couponId)

	if err != nil {
		return
	}

	defer rows.Close()

	if rows.Next() {
		var coupon models.Coupon

		err = rows.Scan(&coupon.ID, &coupon.Name, &coupon.CreatedAt, &coupon.UpdatedAt)

		if err != nil {
			return
		}

		result = coupon
	} else {
		err = fmt.Errorf("there is no coupon found with this identifier (%d)", couponId)
	}

	return
}

func InsertCoupon(db *sql.DB, coupon models.CreateCoupon) (err error) {
	sqlStmt := "INSERT INTO coupons (name, discount, created_at, updated_at) VALUES " +
		"($1, $2, $3, $4)"

	disc := coupon.Discount

	if disc < 0 {
		disc = 0
	} else if disc > 100 {
		disc = 100
	}

	errs := db.QueryRow(sqlStmt, coupon.Name, disc, time.Now(), time.Now())

	return errs.Err()
}

func UpdateCoupon(db *sql.DB, coupon models.Coupon) (err error) {
	sqlStatement := "UPDATE coupons SET name=$1, updated_at=$2, " +
		"WHERE id=$3 Returning id"

	errs := db.QueryRow(sqlStatement, coupon.Name,
		time.Now(), coupon.ID).Scan(&coupon.ID)

	return errs
}

func DeleteCoupon(db *sql.DB, couponId int) (err error) {
	sql := "DELETE FROM coupons WHERE id=$1 Returning id"

	errs := db.QueryRow(sql, couponId).Scan(&couponId)

	return errs
}
