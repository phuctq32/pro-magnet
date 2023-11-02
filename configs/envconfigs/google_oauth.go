package envconfigs

type GoogleOAuthConfig interface {
	ClientId() string
	ClientSecret() string
	RedirectUri() string
}

type googleOauth struct {
	env *envConfigs
}

func (ggOauth *googleOauth) ClientId() string {
	return ggOauth.env.GoogleOauthClientId
}

func (ggOauth *googleOauth) ClientSecret() string {
	return ggOauth.env.GoogleOauthClientSecret
}

func (ggOauth *googleOauth) RedirectUri() string {
	return ggOauth.env.GoogleOauthRedirectUri
}
