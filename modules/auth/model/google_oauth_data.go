package authmodel

import "time"

type GoogleOAuthData struct {
	Url             string
	State           string
	StateExpiration time.Duration
}
