package command

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/sparhokm/go-course-ms-auth/pkg/auth_v1"
)

type user struct {
	m                   sync.RWMutex
	accessToken         string
	refreshToken        string
	email               string
	authClient          auth_v1.AuthV1Client
	ctx                 context.Context
	accessTokenRefresh  time.Duration
	refreshTokenRefresh time.Duration
}

func NewUser(ctx context.Context, authClient auth_v1.AuthV1Client, email string, refreshToken string) *user {
	u := &user{
		refreshToken:        refreshToken,
		authClient:          authClient,
		email:               email,
		ctx:                 ctx,
		accessTokenRefresh:  5 * time.Second,
		refreshTokenRefresh: 20 * time.Second,
	}
	u.updateAccessToken()

	go func() {
		accessTokenTicker := time.Tick(u.accessTokenRefresh)

		for {
			select {
			case <-accessTokenTicker:
				u.updateAccessToken()
			case <-ctx.Done():
				fmt.Println("Done")
				return
			}
		}
	}()
	go func() {
		refreshTokenTicker := time.Tick(u.refreshTokenRefresh)

		for {
			select {
			case <-refreshTokenTicker:
				u.updateRefreshToken()
			case <-ctx.Done():
				fmt.Println("Done")
				return
			}
		}
	}()

	return u
}

func (u *user) updateAccessToken() {
	u.m.Lock()
	auth, err := u.authClient.GetAccessToken(u.ctx, &auth_v1.GetAccessTokenIn{RefreshToken: u.refreshToken})
	if err != nil {
		log.Fatalf("Can't update access token: %s \n", err)
	}
	u.accessToken = auth.GetAccessToken()
	u.m.Unlock()
}

func (u *user) updateRefreshToken() {
	u.m.Lock()
	auth, err := u.authClient.GetRefreshToken(u.ctx, &auth_v1.GetRefreshTokenIn{RefreshToken: u.refreshToken})
	if err != nil {
		log.Fatalf("Can't update refresh token: %s \n", err)
	}
	u.refreshToken = auth.GetRefreshToken()
	u.m.Unlock()
}

func (u *user) GetAccessToken() string {
	u.m.RLock()
	token := u.accessToken
	u.m.RUnlock()
	return token
}

func (u *user) GetEmail() string {
	return u.email
}
