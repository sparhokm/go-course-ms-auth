package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/wrapperspb"

	model "github.com/sparhokm/go-course-ms-auth/internal/model/user"
	"github.com/sparhokm/go-course-ms-auth/internal/rpc/user"
	userServiceMock "github.com/sparhokm/go-course-ms-auth/internal/rpc/user/mocks"
	desc "github.com/sparhokm/go-course-ms-auth/pkg/user_v1"
)

func TestUpdate(t *testing.T) {
	t.Parallel()

	var (
		ctx = context.Background()

		serviceErr = fmt.Errorf("service error")

		name    = "Name"
		email   = "email@email.test"
		role, _ = model.NewRole("user")

		roleRpc  = desc.Role_USER
		nameRpc  = wrapperspb.StringValue{Value: name}
		emailRpc = wrapperspb.StringValue{Value: email}

		id int64 = 100
	)

	tests := []struct {
		name      string
		prepare   func(*userServiceMock.MockUserService)
		updateIn  *desc.UpdateIn
		expectErr bool
	}{
		{
			name: "success case",
			updateIn: &desc.UpdateIn{
				Id:    id,
				Name:  &nameRpc,
				Role:  &roleRpc,
				Email: &emailRpc,
			},
			expectErr: false,
			prepare: func(m *userServiceMock.MockUserService) {
				m.EXPECT().Update(ctx, id, &name, &email, role).Return(nil)
			},
		},
		{
			name: "user service error",
			updateIn: &desc.UpdateIn{
				Id:    id,
				Name:  &nameRpc,
				Role:  &roleRpc,
				Email: &emailRpc,
			},
			expectErr: true,
			prepare: func(m *userServiceMock.MockUserService) {
				m.EXPECT().Update(ctx, id, &name, &email, role).Return(serviceErr)
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

			_, err := api.Update(ctx, tt.updateIn)

			require.Equal(t, tt.expectErr, err != nil)
		})
	}
}
