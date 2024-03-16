package user

import (
	"context"
	"errors"

	"github.com/sparhokm/go-course-ms-auth/internal/model/user"
	"github.com/sparhokm/go-course-ms-auth/pkg/client/db"
)

type service struct {
	userRepo        UserRepo
	userDeletedRepo DeletedUserRepo
	txManager       db.TxManager
	hasher          Hasher
}

func NewUserService(userRepo UserRepo, userDeletedRepo DeletedUserRepo, hasher Hasher, txManager db.TxManager) *service {
	return &service{
		userRepo:        userRepo,
		userDeletedRepo: userDeletedRepo,
		txManager:       txManager,
		hasher:          hasher,
	}
}

func (u service) Create(ctx context.Context, userInfo *user.Info, newPassword *user.Password) (int64, error) {
	userModel, _ := u.userRepo.GetByEmail(ctx, userInfo.Email)
	if userModel != nil {
		return 0, errors.New("email exist")
	}

	passwordHash, err := u.hasher.Hash(newPassword.GetPassword())
	if err != nil {
		return 0, err
	}

	return u.userRepo.Save(ctx, userInfo, passwordHash)
}

func (u service) Delete(ctx context.Context, id int64) error {
	return u.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		err := u.userRepo.Delete(ctx, id)
		if err != nil {
			return err
		}

		return u.userDeletedRepo.Save(ctx, id)
	})
}

func (u service) Get(ctx context.Context, id int64) (*user.User, error) {
	return u.userRepo.Get(ctx, id)
}

func (u service) Update(ctx context.Context, id int64, name *string, email *string, role *user.Role) error {
	return u.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		userModel, err := u.userRepo.Get(ctx, id)
		if err != nil {
			return err
		}

		info := userModel.Info
		needUpdate := updater(&info.Name, name)
		needUpdate = needUpdate || updater(&info.Email, email)
		needUpdate = needUpdate || updater(&info.Role, role)

		if !needUpdate {
			return nil
		}

		return u.userRepo.Update(ctx, id, &info)
	})
}

func updater[T comparable](v *T, newVal *T) bool {
	if newVal != nil && *newVal != *v {
		*v = *newVal
		return true
	}

	return false
}
