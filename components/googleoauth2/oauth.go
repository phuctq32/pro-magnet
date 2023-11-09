package ggoauth2

import (
	"context"
	ggoauthmodel "pro-magnet/components/googleoauth2/model"
)

type GoogleOAuth interface {
	AuthURL(state string) string
	UserInfoFromCode(ctx context.Context, code string) (*ggoauthmodel.User, error)
}
