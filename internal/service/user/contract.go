package user

import (
	"context"

	"github.com/sparhokm/go-course-ms-auth/internal/model/user"
)

type UserRepo interface {
	Save(ctx context.Context, info *user.Info, passwordHash string) (int64, error)
	Delete(ctx context.Context, id int64) error
	Get(ctx context.Context, id int64) (*user.User, error)
	Update(ctx context.Context, id int64, info *user.Info) error
}

type DeletedUserRepo interface {
	Save(ctx context.Context, id int64) error
}

type Hasher interface {
	Hash(password string) (string, error)
}
