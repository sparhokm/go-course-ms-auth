package user

import (
	"strings"

	"google.golang.org/protobuf/types/known/timestamppb"

	model "github.com/sparhokm/go-course-ms-auth/internal/model/user"
	desc "github.com/sparhokm/go-course-ms-auth/pkg/user_v1"
)

func ToUserRpc(user *model.User) *desc.GetOut {
	var updateAt *timestamppb.Timestamp
	if user.UpdatedAt != nil {
		updateAt = timestamppb.New(*user.UpdatedAt)
	}

	return &desc.GetOut{
		Id:        user.Id,
		UserInfo:  ToUserInfoRpc(&user.Info),
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: updateAt,
	}
}

func ToRoleRpc(role model.Role) desc.Role {
	return desc.Role(desc.Role_value[strings.ToUpper(role.GetCode())])
}

func ToUserInfoRpc(info *model.Info) *desc.UserInfo {
	return &desc.UserInfo{
		Name:  info.Name,
		Email: info.Email,
		Role:  ToRoleRpc(info.Role),
	}
}
