package repositories

import (
	"database/sql"
	"fmt"
	"sb-pos/models"
	"time"
)

func GetAllUsers(db *sql.DB) (results []models.User, err error) {
	sqlStmt := "SELECT u.id, u.username, u.email, u.created_at, u.updated_at, " +
		"u.role_id, r.name FROM users u " +
		"JOIN roles r ON u.role_id = r.id ORDER BY u.created_at DESC"

	rows, err := db.Query(sqlStmt)
	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		var user = models.User{}

		err = rows.Scan(&user.ID, &user.Username, &user.Email, &user.CreatedAt,
			&user.UpdatedAt, &user.RoleId, &user.RoleName)
		if err != nil {
			return
		}

		results = append(results, user)
	}

	return
}

func GetUserById(db *sql.DB, userId int) (result models.User, err error) {
	sqlStatement := "SELECT u.id, u.username, u.email, u.created_at, u.updated_at, " +
		"u.role_id, r.name FROM users u " +
		"JOIN roles r ON u.role_id = r.id WHERE u.id=$1"

	rows, errQuery := db.Query(sqlStatement, userId)
	if errQuery != nil {
		err = errQuery
		return
	}

	if rows.Next() {
		var user = models.User{}

		err = rows.Scan(&user.ID, &user.Username, &user.Email, &user.CreatedAt,
			&user.UpdatedAt, &user.RoleId, &user.RoleName)
		if err != nil {
			return
		}

		result = user
	} else {
		err = fmt.Errorf("there is no user found with this identifier (%d)", userId)
	}

	return
}

func UpdateUser(db *sql.DB, user models.UpdateUser) (err error) {
	sqlStatement := "UPDATE users SET username=$1, email=$2, " +
		"updated_at=$3 WHERE id=$4 Returning id"

	errs := db.QueryRow(sqlStatement, user.Username, user.Email,
		time.Now(), user.ID).Scan(&user.ID)

	return errs
}

func UpdateUserRole(db *sql.DB, user models.UpdateUserRole) (err error) {
	sqlStatement := "UPDATE users SET role_id=$1, " +
		"updated_at=$2 WHERE id=$3 Returning id"

	errs := db.QueryRow(sqlStatement, user.RoleId, time.Now(), user.ID).Scan(&user.ID)

	return errs
}

func DeleteUser(db *sql.DB, userId int) (err error) {
	sql := "DELETE FROM users WHERE id=$1 Returning id"

	errs := db.QueryRow(sql, userId).Scan(&userId)

	return errs
}
