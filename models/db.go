package models

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var db *sql.DB

// InitDB initializes the PostgreSQL database connection
func InitDB() {
	connStr := "user=yourusername dbname=yourdbname sslmode=disable password=yourpassword"
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("Failed to connect to the database")
		panic(err)
	}
}
