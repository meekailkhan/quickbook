package utils

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

func ConnectAndUpDB() (*sql.DB, error) {
	databaseUrl := os.Getenv("ENV_DATABASE_URL")
	if databaseUrl == "" {
		return nil, errors.New("Database Url not set in enviourment variable")
	}
	conn, err := sql.Open("pgx", databaseUrl)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	if err := conn.Ping(); err != nil {
		fmt.Println(err)
		return nil, errors.New("Failed to reached database and ping it")
	}

	if err := goose.SetDialect("postgres"); err != nil {
		fmt.Printf("set goose dialect: %v\n", err)
		return nil, errors.New("set goose dialect error")
	}

	if err := goose.Up(conn, "migrations"); err != nil {
		fmt.Printf("run migrations: %v\n", err)
		return nil, errors.New("run migrations error")
	}
	fmt.Println("***successfully connected to the db and migrate***")
	return conn, nil

}
