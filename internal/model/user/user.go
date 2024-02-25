package user

import (
	"time"
)

type User struct {
	Id           int64
	Info         Info
	PasswordHash string
	CreatedAt    time.Time
	UpdatedAt    *time.Time
}
