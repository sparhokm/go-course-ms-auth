package test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"

	"github.com/sparhokm/go-course-ms-auth/internal/service/hasher"
)

func TestSuccess(t *testing.T) {
	t.Parallel()

	password := "123456"
	h := hasher.NewHasher()
	hash, err := h.Hash(password)

	require.NoError(t, bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)))
	require.NoError(t, err)
}

func TestWrong(t *testing.T) {
	t.Parallel()

	password := "123456"
	h := hasher.NewHasher()
	hash, err := h.Hash(password)

	require.Error(t, bcrypt.CompareHashAndPassword([]byte(hash), []byte("123")))
	require.NoError(t, err)
}
