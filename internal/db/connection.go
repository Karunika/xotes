package db

import (
	"database/sql"
	"fmt"

	"server/internal/config"

	_ "github.com/lib/pq"
)

type DBService struct {
	db *sql.DB
}

func GenerateDBURL() string {
	dbConfig := config.GetDatabaseConfig()
	dbURL := fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Name,
		dbConfig.User,
		dbConfig.Password,
	)
	return dbURL
}

func ConnectDB(dbURL string) (*DBService, error) {
	db, err := sql.Open("postgres", dbURL)

	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %w", err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return &DBService{db}, nil
}

func (thisService *DBService) CloseDB() error {
	if thisService.db == nil {
		return nil
	}

	return thisService.db.Close()
}
