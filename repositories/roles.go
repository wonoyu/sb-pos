package repositories

import (
	"database/sql"
	"fmt"
	"sb-pos/models"
	"time"
)

func GetAllRoles(db *sql.DB) (results []models.Role, err error) {
	sqlStmt := "SELECT * FROM roles ORDER BY created_at DESC"

	rows, err := db.Query(sqlStmt)
	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		var role = models.Role{}

		err = rows.Scan(&role.ID, &role.Name, &role.CreatedAt, &role.UpdatedAt)
		if err != nil {
			return
		}

		results = append(results, role)
	}

	return
}

func GetRoleById(db *sql.DB, roleId int) (result models.Role, err error) {
	sqlStatement := "SELECT * FROM roles WHERE id=$1"

	rows, errQuery := db.Query(sqlStatement, roleId)
	if errQuery != nil {
		err = errQuery
		return
	}

	if rows.Next() {
		var role = models.Role{}

		err = rows.Scan(&role.ID, &role.Name, &role.CreatedAt, &role.UpdatedAt)
		if err != nil {
			return
		}

		result = role
	} else {
		err = fmt.Errorf("there is no role found with this identifier (%d)", roleId)
	}

	return
}

func InsertRole(db *sql.DB, role models.Role) (err error) {
	sql := "INSERT INTO roles (name, created_at, updated_at)" +
		"VALUES ($1, $2, $3)"

	errs := db.QueryRow(sql, &role.Name, time.Now(), time.Now())

	return errs.Err()
}

func UpdateRole(db *sql.DB, role models.Role) (err error) {
	sqlStatement := "UPDATE roles SET name=$1, updated_at=$2 " +
		"WHERE id=$3 Returning id"

	errs := db.QueryRow(sqlStatement, role.Name,
		time.Now(), role.ID).Scan(&role.ID)

	return errs
}

func DeleteRole(db *sql.DB, roleId int) (err error) {
	sql := "DELETE FROM roles WHERE id=$1 Returning id"

	errs := db.QueryRow(sql, roleId).Scan(&roleId)

	return errs
}
