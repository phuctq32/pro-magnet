package envconfigs

type CloudinaryConfig interface {
	CloudName() string
	ApiKey() string
	ApiSecret() string
}

type cloudinaryConfig struct {
	env *envConfigs
}

func (cfg *cloudinaryConfig) CloudName() string {
	return cfg.env.CldCloudName
}

func (cfg *cloudinaryConfig) ApiKey() string {
	return cfg.env.CldApiKey
}

func (cfg *cloudinaryConfig) ApiSecret() string {
	return cfg.env.CldApiSecret
}
