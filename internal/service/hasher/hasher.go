package hasher

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type hasher struct{}

func NewHasher() *hasher {
	return &hasher{}
}

func (h hasher) Hash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("error generate hash: %w", err)
	}

	return string(hash), nil
}

func (h hasher) Verify(hash string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
