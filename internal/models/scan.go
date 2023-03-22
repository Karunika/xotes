package models

import (
	"database/sql"
)

func ScanUser(u *User, rows *sql.Rows) error {
	if err := rows.Scan(
		&u.Uuid,
		&u.Email,
		&u.Username,
		&u.Country,
		&u.Bio,
		&u.BfpBlob,
		&u.PfpMimeType,
		&u.DateCreated,
	); err != nil {
		return err
	}
	return nil
}

func ScanUserCred(u *Auth, rows *sql.Rows) error {
	if err := rows.Scan(
		&u.Email,
		&u.Username,
		&u.Password,
	); err != nil {
		return err
	}
	return nil
}
