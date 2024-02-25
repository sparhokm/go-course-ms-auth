package user

import (
	"errors"
)

var (
	ErrNameRequired  = errors.New("name required")
	ErrEmailRequired = errors.New("email required")
)

type Info struct {
	Name  string
	Email string
	Role  Role
}

func NewUserInfo(name string, email string, role Role) (*Info, error) {
	if name == "" {
		return nil, ErrNameRequired
	}

	if email == "" {
		return nil, ErrEmailRequired
	}

	return &Info{name, email, role}, nil
}
