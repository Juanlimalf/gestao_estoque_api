package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	DB_HOST     = "localhost"
	DB_PORT     = 5432
	DB_USER     = "db_user"
	DB_PASSWORD = "senha123"
	DB_NAME     = "db_test"
)

func Connect() (*sql.DB, error) {

	InfoDatabase := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME)

	conn, err := sql.Open("postgres", InfoDatabase)

	if err != nil {
		fmt.Print("Could not connect to the database: ", err)
		return nil, err
	}

	fmt.Println("Connected to the database successfully")

	return conn, nil
}
