package tokenGenerator

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/sparhokm/go-course-ms-auth/internal/model/user"
)

type UserClaims struct {
	jwt.RegisteredClaims
	Id   int64     `json:"i"`
	Role user.Role `json:"r"`
}
