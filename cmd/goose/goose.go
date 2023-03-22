package main

import (
	"server/internal/db"
)

func main() {
	dbURL := db.GenerateDBURL()
	dbService, _ := db.ConnectDB(dbURL)
	dbService.RunMigrations()
	defer dbService.CloseDB()
}
