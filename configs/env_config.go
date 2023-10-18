package configs

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"reflect"
)

const (
	Development string = "development"
	Production         = "production"
)

var EnvConfigs EnvConfiguration

type EnvConfiguration interface {
	Port() int
	MongoConnectionString() string
	MongoDBName() string
	RedisAddr() string
	SendgridApiKey() string
	SendgridFromEmail() string
	SendgridVerifyEmailTemplateId() string
	VerificationURL() string
	AccessSecret() string
	RefreshSecret() string
	AccessTokenExpiry() int
	RefreshTokenExpiry() int
	S3BucketName() string
	S3AccessKey() string
	S3SecretKey() string
	S3Region() string
	S3Domain() string
	CloudinaryCloudName() string
	CloudinaryApiKey() string
	CloudinaryApiSecret() string
}

type envConfigs struct {
	env struct {
		// App
		Port               int    `mapstructure:"PORT"`
		AccessSecret       string `mapstructure:"ACCESS_SECRET"`
		RefreshSecret      string `mapstructure:"REFRESH_SECRET"`
		AccessTokenExpiry  int    `mapstructure:"ACCESS_TOKEN_EXPIRY"`
		RefreshTokenExpiry int    `mapstructure:"REFRESH_TOKEN_EXPIRY"`

		// MongoDB
		MongoConnStr string `mapstructure:"MONGO_CONN_STRING"`
		MongoDBName  string `mapstructure:"MONGO_DATABASE_NAME"`

		// Redis
		RedisAddr string `mapstructure:"REDIS_ADDR"`

		// Sendgrid
		SendGridApiKey                string `mapstructure:"SENDGRID_API_KEY"`
		SendGridFromEmail             string `mapstructure:"FROM_EMAIL"`
		SendGridVerifyEmailTemplateId string `mapstructure:"SENDGRID_VERIFY_TEMPlATE_ID"`
		VerificationURL               string `mapstructure:"VERIFICATION_URL"`

		// AWS S3
		S3BucketName string `mapstructure:"S3_BUCKET_NAME"`
		S3AccessKey  string `mapstructure:"S3_ACCESS_KEY"`
		S3SecretKey  string `mapstructure:"S3_SECRET_KEY"`
		S3Region     string `mapstructure:"S3_REGION"`
		S3Domain     string `mapstructure:"S3_DOMAIN"`

		// Cloudinary
		CldApiKey    string `mapstructure:"CLOUDINARY_API_KEY"`
		CldApiSecret string `mapstructure:"CLOUDINARY_API_SECRET"`
		CldCloudName string `mapstructure:"CLOUDINARY_CLOUD_NAME"`
	}
}

func LoadEnvConfigs(env string) {
	cfg := new(envConfigs)

	if env == Development {
		cfg.LoadFromEnvFile()
	} else if env == Production {
		cfg.LoadConfigsFromOS()
	} else {
		log.Fatal().Msg("invalid environment")
	}

	EnvConfigs = cfg
}

func (cfg *envConfigs) LoadFromEnvFile() {
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal().Err(err).Msg("cannot load environment variables from env file")
	}

	if err := viper.Unmarshal(&cfg.env); err != nil {
		log.Fatal().Err(err).Msg("unmarshal env error happened")
	}
}

func (cfg *envConfigs) LoadConfigsFromOS() {
	t := reflect.ValueOf(&cfg.env).Elem().Type()
	for i := 0; i < t.NumField(); i++ {
		if tagValue := t.Field(i).Tag.Get("mapstructure"); tagValue != "-" {
			if err := viper.BindEnv(tagValue); err != nil {
				log.Fatal().Err(err).Msgf("cannot bind env: %v", tagValue)
			}
		}
	}

	if err := viper.Unmarshal(&cfg.env); err != nil {
		log.Fatal().Err(err).Msg("unmarshal env error happened")
	}
}

func (cfg *envConfigs) Port() int {
	return cfg.env.Port
}

func (cfg *envConfigs) MongoConnectionString() string {
	return cfg.env.MongoConnStr
}

func (cfg *envConfigs) MongoDBName() string {
	return cfg.env.MongoDBName
}

func (cfg *envConfigs) RedisAddr() string {
	return cfg.env.RedisAddr
}

func (cfg *envConfigs) SendgridApiKey() string {
	return cfg.env.SendGridApiKey
}

func (cfg *envConfigs) SendgridFromEmail() string {
	return cfg.env.SendGridFromEmail
}

func (cfg *envConfigs) SendgridVerifyEmailTemplateId() string {
	return cfg.env.SendGridVerifyEmailTemplateId
}

func (cfg *envConfigs) VerificationURL() string {
	return cfg.env.VerificationURL
}

func (cfg *envConfigs) AccessSecret() string {
	return cfg.env.AccessSecret
}

func (cfg *envConfigs) RefreshSecret() string {
	return cfg.env.RefreshSecret
}

func (cfg *envConfigs) AccessTokenExpiry() int {
	return cfg.env.AccessTokenExpiry
}

func (cfg *envConfigs) RefreshTokenExpiry() int {
	return cfg.env.RefreshTokenExpiry
}

func (cfg *envConfigs) S3BucketName() string {
	return cfg.env.S3BucketName
}

func (cfg *envConfigs) S3AccessKey() string {
	return cfg.env.S3AccessKey
}

func (cfg *envConfigs) S3SecretKey() string {
	return cfg.env.S3SecretKey
}

func (cfg *envConfigs) S3Region() string {
	return cfg.env.S3Region
}

func (cfg *envConfigs) S3Domain() string {
	return cfg.env.S3Domain
}

func (cfg *envConfigs) CloudinaryCloudName() string {
	return cfg.env.CldCloudName
}

func (cfg *envConfigs) CloudinaryApiKey() string {
	return cfg.env.CldApiKey
}

func (cfg *envConfigs) CloudinaryApiSecret() string {
	return cfg.env.CldApiSecret
}
