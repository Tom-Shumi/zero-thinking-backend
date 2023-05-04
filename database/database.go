package database

import (
	"database/sql"
	"fmt"
	"zero-thinking-backend/config"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// Init initializes database
func Init() {
	c := config.GetConfig()
	var err error
	db, err = sql.Open(c.GetString("db.provider"),
		fmt.Sprintf("%s:%s@(%s:%s)/zeroThinking?parseTime=true", c.GetString("db.user"), c.GetString("db.password"), c.GetString("db.domain"), c.GetString("db.port")))
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
