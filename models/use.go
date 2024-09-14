package models

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type User struct {
	ID       string
	Email    string
	Password string
}

// CreateUser inserts a new user into the database
func CreateUser(user *User) error {
	_, err := db.Exec("INSERT INTO users (id, email, password) VALUES ($1, $2, $3)", user.ID, user.Email, user.Password)
	return err
}

// GetUserByEmail retrieves a user by email
func GetUserByEmail(email string) (*User, error) {
	row := db.QueryRow("SELECT id, email, password FROM users WHERE email = $1", email)

	var user User
	err := row.Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no user found with that email")
		}
		return nil, err
	}

	return &user, nil
}
