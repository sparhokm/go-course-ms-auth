package user

import (
	"errors"
)

var (
	ErrWrongConfirm     = errors.New("wrong confirm password")
	ErrPasswordVeryEasy = errors.New("password very easy")
)

type Password struct {
	value string
}

func NewPassword(password string, confirm string) (*Password, error) {
	if len(password) <= 6 {
		return nil, ErrPasswordVeryEasy
	}

	if password != confirm {
		return nil, ErrWrongConfirm
	}

	return &Password{password}, nil
}

func (p Password) GetPassword() string {
	return p.value
}
