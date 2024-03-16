package access

import (
	"context"

	"github.com/sparhokm/go-course-ms-auth/internal/model/user"
	"github.com/sparhokm/go-course-ms-auth/internal/service/tokenGenerator"
)

type TokenVerifier interface {
	VerifyToken(tokenStr string) (*tokenGenerator.UserClaims, error)
}

type AccessRepo interface {
	Get(ctx context.Context, endpoint string) ([]user.Role, error)
}
