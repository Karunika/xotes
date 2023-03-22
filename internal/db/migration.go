package db

import (
	"embed"
	"log"

	"github.com/pressly/goose/v3"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

func (thisService *DBService) RunMigrations() {
	goose.SetDialect("postgres")
	goose.SetBaseFS(embedMigrations)
	if err := goose.Down(thisService.db, "migrations"); err != nil {
		panic(err)
	}
	if err := goose.Up(thisService.db, "migrations"); err != nil {
		panic(err)
	}
	if err := goose.Version(thisService.db, "migrations"); err != nil {
		log.Fatal(err)
	}
}
