package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"rutube/config"
)

var DB *sql.DB

func InitDB() {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Cfg.DBHost, config.Cfg.DBPort, config.Cfg.DBUser, config.Cfg.DBPassword, config.Cfg.DBName)
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error connecting to the database: %s", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatalf("Error pinging the database: %s", err)
	}
}
