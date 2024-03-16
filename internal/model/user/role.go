package user

import (
	"errors"
	"strings"
)

var (
	ErrUnknownRole = errors.New("unknown role")
)

var availableRoles = map[string]struct{}{
	"user":  {},
	"admin": {},
}

type Role string

func NewRole(code string) (*Role, error) {
	code = strings.ToLower(code)

	if _, ok := availableRoles[code]; !ok {
		return nil, ErrUnknownRole
	}

	r := Role(code)
	return &r, nil
}

func (r Role) GetCode() string {
	return string(r)
}

func (r Role) EqualTo(role *Role) bool {
	return role.GetCode() == r.GetCode()
}
