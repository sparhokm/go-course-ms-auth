package access

import (
	"context"
	"errors"
)

type service struct {
	accessTokenVerifier TokenVerifier
	accessRepo          AccessRepo
}

func NewAccessService(accessTokenVerifier TokenVerifier, accessRepo AccessRepo) *service {
	return &service{
		accessTokenVerifier: accessTokenVerifier,
		accessRepo:          accessRepo,
	}
}

func (a service) Check(ctx context.Context, token string, endpoint string) error {
	userClaims, err := a.accessTokenVerifier.VerifyToken(token)
	if err != nil {
		return err
	}

	roles, err := a.accessRepo.Get(ctx, endpoint)
	if err != nil {
		return err
	}

	for _, r := range roles {
		if r.EqualTo(&userClaims.Role) {
			return nil
		}
	}

	return errors.New("access denied")
}
