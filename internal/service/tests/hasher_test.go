package tests

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"

	"github.com/sparhokm/go-course-ms-auth/internal/service"
)

func TestSuccess(t *testing.T) {
	t.Parallel()

	password := "123456"
	hasher := service.NewHasher()
	hash, err := hasher.Hash(password)

	require.NoError(t, bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)))
	require.NoError(t, err)
}

func TestWrong(t *testing.T) {
	t.Parallel()

	password := "123456"
	hasher := service.NewHasher()
	hash, err := hasher.Hash(password)

	require.Error(t, bcrypt.CompareHashAndPassword([]byte(hash), []byte("123")))
	require.NoError(t, err)
}
