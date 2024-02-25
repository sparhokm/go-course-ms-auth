package user

import (
	"time"

	model "github.com/sparhokm/go-course-ms-auth/internal/model/user"
	repoModel "github.com/sparhokm/go-course-ms-auth/internal/repository/user/model"
)

func ToUserFromRepo(user *repoModel.User) (*model.User, error) {
	role, err := model.NewRole(user.Info.Role)
	if err != nil {
		return nil, err
	}

	var updatedAt *time.Time
	if user.UpdatedAt.Valid {
		updatedAt = &user.UpdatedAt.Time
	}

	return &model.User{
		Id: user.Id,
		Info: model.Info{
			Name:  user.Info.Name,
			Email: user.Info.Email,
			Role:  *role,
		},
		CreatedAt: user.CreatedAt,
		UpdatedAt: updatedAt,
	}, nil
}
