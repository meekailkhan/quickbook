package util

import (
	"database/sql"
	"errors"
	"os"

	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {
	databaseUrl := os.Getenv("ENV_DATABASE_URL")
	if databaseUrl == "" {
		return nil, errors.New("Database Url not set in enviourment variable")
	}
	conn, err := sql.Open("postgres", databaseUrl)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	if err := conn.Ping(); err != nil {
		return nil, errors.New("Failed to reached database and ping it")
	}
	return conn, nil

}
