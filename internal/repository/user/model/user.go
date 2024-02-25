package model

import (
	"database/sql"
	"time"
)

type User struct {
	Id           int64
	Info         Info `db:""`
	PasswordHash string
	CreatedAt    time.Time
	UpdatedAt    sql.NullTime
}

type Info struct {
	Name  string
	Email string
	Role  string
}
