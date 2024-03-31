package command

import (
	"context"
	"strconv"

	"github.com/sparhokm/go-course-ms-auth/pkg/user_v1"
)

type userNames struct {
	names      map[int64]string
	userClient user_v1.UserV1Client
}

func NewUserNames(userClient user_v1.UserV1Client) *userNames {
	return &userNames{userClient: userClient, names: make(map[int64]string)}
}

func (u *userNames) GetName(ctx context.Context, id int64) string {
	if email, ok := u.names[id]; ok {
		return email
	}

	var name string
	userData, err := u.userClient.Get(ctx, &user_v1.GetIn{Id: id})
	if err != nil {
		name = strconv.FormatInt(id, 10)
	} else {
		name = userData.GetUserInfo().GetName()
	}
	u.names[id] = name

	return name
}
