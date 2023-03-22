package db

import (
	"database/sql"
	"fmt"
	"log"

	"server/internal/models"
)

type QueryFunc[T models.Models] func(string, ...any) []T

func Select[T models.Models](thisService *DBService, callback func(*sql.Rows) []T) QueryFunc[T] {
	return func(query string, args ...any) []T {
		rows, err := thisService.db.Query(query, args...)
		if err != nil {
			log.Fatal(err)
		}
		return callback(rows)
	}
}

func RowsToSlice[T models.Models](ScanFunc func(*T, *sql.Rows) error) func(*sql.Rows) []T {
	return func(rows *sql.Rows) []T {
		var objs []T
		for rows.Next() {
			var obj T
			ScanFunc(&obj, rows)
			objs = append(objs, obj)
			fmt.Println(obj)
		}
		return objs
	}
}

func (thisService *DBService) SelectUsers() QueryFunc[models.User] {
	return Select(thisService, RowsToSlice(models.ScanUser))
}

func (thisService *DBService) SelectUserCred() QueryFunc[models.Auth] {
	return Select(thisService, RowsToSlice(models.ScanUserCred))
}

func (thisService *DBService) SelectAllUsers() []models.User {
	return thisService.SelectUsers()(`SELECT * FROM users`)
}

func (thisService *DBService) SelectUserByUsername(username string) models.User {
	query := `SELECT uuid, email, username, country, bio, pfp_blob, pfp_mime_type, date_created FROM users WHERE users.username = $1`
	user := thisService.SelectUsers()(query, username)
	if len(user) == 1 {
		return user[0]
	}
	panic("User not found")
}

func (thisService *DBService) SelectUserCredByKey(key string) models.Auth {
	query := `SELECT email, username, pwd_hash FROM users WHERE users.username = $1 OR users.email = $1`
	user := thisService.SelectUserCred()(query, key)
	if len(user) == 1 {
		return user[0]
	}
	panic("User not found")
}

func (thisService *DBService) InsertUser(username string, email string, pwdHash string) {
	query := `INSERT INTO users(username, email, pwd_hash, country) VALUES ($1, $2, $3, 'Algeria')`
	_, err := thisService.db.Exec(query, username, email, pwdHash)
	if err != nil {
		panic(err)
	}
}
