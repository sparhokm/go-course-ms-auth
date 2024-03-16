package user

import (
	"context"

	"github.com/sparhokm/go-course-ms-auth/internal/model/user"
)

type UserService interface {
	Create(ctx context.Context, userInfo *user.Info, newPassword *user.Password) (int64, error)
	Get(ctx context.Context, id int64) (*user.User, error)
	Update(ctx context.Context, id int64, name *string, email *string, role *user.Role) error
	Delete(ctx context.Context, id int64) error
}

type TokenGenerator interface {
	GenerateToken(id int64, role user.Role) (string, error)
}
