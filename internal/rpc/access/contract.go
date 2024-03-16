package access

import (
	"context"
)

type AccessService interface {
	Check(ctx context.Context, accessToken string, path string) error
}
