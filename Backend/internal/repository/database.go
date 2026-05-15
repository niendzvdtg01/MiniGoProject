package respository

import (
	"database/sql"
)

var DB *sql.DB

func ConnectDB() error {
	var err error
	db, err := sql.Open("mysql", "user:password@tcp(localhost)/dbname")
	if err != nil {
		return err
	}
	DB = db
	return nil
}
