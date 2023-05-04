package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// Init initializes database
func Init() {
	// c := config.GetConfig()
	var err error
	db, err = sql.Open("mysql", "root:root@(localhost:3306)/zeroThinking?parseTime=true")
	if err != nil {
		panic(err)
	}
}

// GetDB returns database connection
func GetDB() *sql.DB {
	return db
}

// Close closes database
func Close() {
	db.Close()
}
