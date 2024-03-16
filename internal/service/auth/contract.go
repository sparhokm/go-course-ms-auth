package auth

import (
	"context"

	"github.com/sparhokm/go-course-ms-auth/internal/model/user"
	"github.com/sparhokm/go-course-ms-auth/internal/service/tokenGenerator"
)

type UserRepo interface {
	GetByEmail(ctx context.Context, username string) (*user.User, error)
	Get(ctx context.Context, id int64) (*user.User, error)
}

type Hasher interface {
	Verify(hash string, password string) error
}

type TokenGenerator interface {
	GenerateToken(id int64, role user.Role) (string, error)
	VerifyToken(tokenStr string) (*tokenGenerator.UserClaims, error)
}
