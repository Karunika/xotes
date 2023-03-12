package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)


// dbURL := fmt.SPrintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable", host, port, name, user, password)

func ConnectDB(dbURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return db, nil
}

// func CloseDB() error {
// 	if DB = nil {
// 		return DB
// 	}

// 	return DB.close()
// }