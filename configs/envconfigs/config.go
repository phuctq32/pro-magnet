package envconfigs

const (
	Development string = "development"
	Production         = "production"
)

type Configs interface {
	App() AppConfig
	AwsS3() S3Config
	Cloudinary() CloudinaryConfig
	Mongo() MongoConfig
	Redis() RedisConfig
	Sendgrid() SendgridConfig
}

type configs struct {
	appCfg   AppConfig
	s3Cfg    S3Config
	cldCfg   CloudinaryConfig
	mongoCfg MongoConfig
	redisCfg RedisConfig
	sgCfg    SendgridConfig
}

func New(env string) Configs {
	envCfg := LoadEnvConfigs(env)

	return &configs{
		appCfg:   &appConfig{env: envCfg},
		s3Cfg:    &s3Config{env: envCfg},
		cldCfg:   &cloudinaryConfig{env: envCfg},
		mongoCfg: &mongoConfig{env: envCfg},
		redisCfg: &redisConfig{env: envCfg},
		sgCfg:    &sendgridConfig{env: envCfg},
	}
}

func (cfg *configs) App() AppConfig {
	return cfg.appCfg
}

func (cfg *configs) AwsS3() S3Config {
	return cfg.s3Cfg
}

func (cfg *configs) Cloudinary() CloudinaryConfig {
	return cfg.cldCfg
}

func (cfg *configs) Mongo() MongoConfig {
	return cfg.mongoCfg
}

func (cfg *configs) Redis() RedisConfig {
	return cfg.redisCfg
}

func (cfg *configs) Sendgrid() SendgridConfig {
	return cfg.sgCfg
}
