package envconfigs

type AppConfig interface {
	Port() int
	AccessSecret() string
	RefreshSecret() string
	AccessTokenExpiry() int
	RefreshTokenExpiry() int
}

type appConfig struct {
	env *envConfigs
}

func (cfg *appConfig) Port() int {
	return cfg.env.Port
}

func (cfg *appConfig) AccessSecret() string {
	return cfg.env.AccessSecret
}

func (cfg *appConfig) RefreshSecret() string {
	return cfg.env.RefreshSecret
}

func (cfg *appConfig) AccessTokenExpiry() int {
	return cfg.env.AccessTokenExpiry
}

func (cfg *appConfig) RefreshTokenExpiry() int {
	return cfg.env.RefreshTokenExpiry
}
