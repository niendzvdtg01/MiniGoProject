package db

import (
	"backend/internal/config"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() error {
	connStr := config.NewConfig().DSN()

	var err error
	DB, err = sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal("Unable to use data source name", err)
	}
	defer DB.Close()

	return err
}
