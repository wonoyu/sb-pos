package repositories

import (
	"database/sql"
	"fmt"
	"sb-pos/models"
	"time"
)

func RegisterUser(db *sql.DB, user models.RegisterUser) (err error) {
	sql := "INSERT INTO users (username, email, password, created_at, updated_at, role_id)" +
		"VALUES ($1, $2, $3, $4, $5, $6)"

	errs := db.QueryRow(sql, &user.Username, &user.Email, &user.Password, time.Now(),
		time.Now(), &user.RoleId)

	return errs.Err()
}

func GetUserByEmail(db *sql.DB, email string) (result models.UserWithPassword, err error) {
	sqlStmt := "SELECT users.*, roles.name FROM users " +
		"JOIN roles ON users.role_id = roles.id WHERE users.email=$1"

	rows, err := db.Query(sqlStmt, email)

	if err != nil {
		return
	}

	defer rows.Close()

	if rows.Next() {
		var user models.UserWithPassword
		err = rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt,
			&user.UpdatedAt, &user.RoleId, &user.RoleName)

		if err != nil {
			return
		}

		result = user
	} else {
		err = fmt.Errorf("user with email %s is not found", email)
		return
	}

	return
}
