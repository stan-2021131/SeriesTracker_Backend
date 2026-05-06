package db

import (
	"database/sql"
	"fmt"
	"os"
)

func ConnectDB() (*sql.DB, error) {
	// String de conexión para producción
	connStr := os.Getenv("DATABASE_URL")

	// String de conexión para desarrollo
	if connStr == "" {
		host := os.Getenv("DB_HOST")
		port := os.Getenv("DB_PORT")
		user := os.Getenv("DB_USER")
		password := os.Getenv("DB_PASSWORD")
		dbname := os.Getenv("DB_NAME")

		connStr = fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			host, port, user, password, dbname,
		)
	}

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func CloseDB(db *sql.DB) {
	db.Close()
}
