package tests

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"

	model "github.com/sparhokm/go-course-ms-auth/internal/model/user"
	"github.com/sparhokm/go-course-ms-auth/internal/rpc/user"
	userServiceMock "github.com/sparhokm/go-course-ms-auth/internal/rpc/user/mocks"
	desc "github.com/sparhokm/go-course-ms-auth/pkg/user_v1"
)

func TestGet(t *testing.T) {
	t.Parallel()

	var (
		ctx = context.Background()

		serviceErr = fmt.Errorf("service error")

		name  = "Name"
		email = "email@email.test"

		id    int64 = 100
		getIn       = &desc.GetIn{
			Id: id,
		}
		createAt  = time.Now()
		updatedAt = time.Now().Add(time.Hour)

		getOut = &desc.GetOut{
			Id: id,
			UserInfo: &desc.UserInfo{
				Name:  name,
				Email: email,
				Role:  desc.Role_USER,
			},
			UpdatedAt: timestamppb.New(updatedAt),
			CreatedAt: timestamppb.New(createAt),
		}
	)

	tests := []struct {
		name      string
		prepare   func(*userServiceMock.MockUserService)
		request   *desc.GetIn
		want      *desc.GetOut
		expectErr bool
	}{
		{
			name:      "success case",
			request:   getIn,
			want:      getOut,
			expectErr: false,
			prepare: func(m *userServiceMock.MockUserService) {
				role, _ := model.NewRole("user")
				m.EXPECT().Get(ctx, id).Return(&model.User{
					Id: id,
					Info: model.Info{
						Name:  name,
						Email: email,
						Role:  *role,
					},
					CreatedAt: createAt,
					UpdatedAt: &updatedAt,
				}, nil)
			},
		},
		{
			name:      "user service error",
			request:   getIn,
			want:      nil,
			expectErr: true,
			prepare: func(m *userServiceMock.MockUserService) {
				m.EXPECT().Get(ctx, id).Return(nil, serviceErr)
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			serviceMock := userServiceMock.NewMockUserService(t)
			tt.prepare(serviceMock)
			api := user.NewImplementation(serviceMock)

			res, err := api.Get(ctx, tt.request)

			require.Equal(t, tt.want, res)
			require.Equal(t, tt.expectErr, err != nil)
		})
	}
}
