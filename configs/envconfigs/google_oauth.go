package envconfigs

type GoogleOAuthConfig interface {
	ClientId() string
	ClientSecret() string
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
