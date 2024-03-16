package tokenGenerator

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/pkg/errors"

	"github.com/sparhokm/go-course-ms-auth/internal/model/user"
)

type tokenGenerator struct {
	secretKey []byte
	duration  time.Duration
}

func NewTokenGenerator(secretKey []byte, duration time.Duration) *tokenGenerator {
	return &tokenGenerator{secretKey, duration}
}

func (t tokenGenerator) GenerateToken(id int64, role user.Role) (string, error) {
	claims := UserClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(t.duration)),
		},
		Id:   id,
		Role: role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(t.secretKey)
}

func (t tokenGenerator) VerifyToken(tokenStr string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(
		tokenStr,
		&UserClaims{},
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, errors.Errorf("unexpected token signing method")
			}

			return t.secretKey, nil
		},
	)
	if err != nil {
		return nil, errors.Errorf("invalid token: %s", err.Error())
	}

	claims, ok := token.Claims.(*UserClaims)
	if !ok {
		return nil, errors.Errorf("invalid token claims")
	}

	return claims, nil
}
