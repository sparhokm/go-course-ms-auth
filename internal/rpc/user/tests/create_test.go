package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/sparhokm/go-course-ms-auth/internal/rpc/user"
	userServiceMock "github.com/sparhokm/go-course-ms-auth/internal/rpc/user/mocks"
	desc "github.com/sparhokm/go-course-ms-auth/pkg/user_v1"
)

func TestCreate(t *testing.T) {
	t.Parallel()

	var (
		ctx = context.Background()

		serviceErr = fmt.Errorf("service error")

		name     = "Name"
		email    = "email@email.test"
		password = "1234567"

		id        int64 = 100
		createOut       = &desc.CreateOut{
			Id: id,
		}
	)

	tests := []struct {
		name      string
		prepare   func(*userServiceMock.MockUserService)
		createIn  *desc.CreateIn
		want      *desc.CreateOut
		expectErr bool
	}{
		{
			name: "success case",
			createIn: &desc.CreateIn{
				UserInfo: &desc.UserInfo{
					Name:  name,
					Role:  desc.Role_USER,
					Email: email,
				},
				Password: &desc.NewPassword{
					Password: password,
					Confirm:  password,
				},
			},
			want:      createOut,
			expectErr: false,
			prepare: func(m *userServiceMock.MockUserService) {
				m.EXPECT().Create(ctx, mock.Anything, mock.Anything).Return(id, nil)
			},
		},
		{
			name: "user service error",
			createIn: &desc.CreateIn{
				UserInfo: &desc.UserInfo{
					Name:  name,
					Role:  desc.Role_ADMIN,
					Email: email,
				},
				Password: &desc.NewPassword{
					Password: password,
					Confirm:  password,
				},
			},
			want:      nil,
			expectErr: true,
			prepare: func(m *userServiceMock.MockUserService) {
				m.EXPECT().Create(ctx, mock.Anything, mock.Anything).Return(0, serviceErr)
			},
		},
		{
			name: "check wrong password",
			createIn: &desc.CreateIn{
				UserInfo: &desc.UserInfo{
					Name:  name,
					Role:  desc.Role_USER,
					Email: email,
				},
				Password: &desc.NewPassword{
					Password: password,
				},
			},
			want:      nil,
			expectErr: true,
			prepare:   func(m *userServiceMock.MockUserService) {},
		},
		{
			name: "check wrong role",
			createIn: &desc.CreateIn{
				UserInfo: &desc.UserInfo{
					Name:  name,
					Role:  desc.Role(213),
					Email: email,
				},
				Password: &desc.NewPassword{
					Password: password,
					Confirm:  password,
				},
			},
			want:      nil,
			expectErr: true,
			prepare:   func(m *userServiceMock.MockUserService) {},
		},
		{
			name: "check empty name",
			createIn: &desc.CreateIn{
				UserInfo: &desc.UserInfo{
					Role:  desc.Role_USER,
					Email: email,
				},
				Password: &desc.NewPassword{
					Password: password,
					Confirm:  password,
				},
			},
			want:      nil,
			expectErr: true,
			prepare:   func(m *userServiceMock.MockUserService) {},
		},
		{
			name: "check empty email",
			createIn: &desc.CreateIn{
				UserInfo: &desc.UserInfo{
					Name: name,
					Role: desc.Role_USER,
				},
				Password: &desc.NewPassword{
					Password: password,
					Confirm:  password,
				},
			},
			want:      nil,
			expectErr: true,
			prepare:   func(m *userServiceMock.MockUserService) {},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			serviceMock := userServiceMock.NewMockUserService(t)
			tt.prepare(serviceMock)
			api := user.NewImplementation(serviceMock)

			res, err := api.Create(ctx, tt.createIn)

			require.Equal(t, tt.want, res)
			require.Equal(t, tt.expectErr, err != nil)
		})
	}
}
