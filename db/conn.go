package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		return nil, err
	}

	// defer db.Close()

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, fmt.Errorf("Error to ping: %w", err)
	}

	fmt.Println("Successfully connected!")

	return db, nil

}
