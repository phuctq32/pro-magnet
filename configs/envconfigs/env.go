package envconfigs

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"reflect"
)

type envConfigs struct {
	// App
	Port               int    `mapstructure:"PORT"`
	AccessSecret       string `mapstructure:"ACCESS_SECRET"`
	RefreshSecret      string `mapstructure:"REFRESH_SECRET"`
	AccessTokenExpiry  int    `mapstructure:"ACCESS_TOKEN_EXPIRY"`
	RefreshTokenExpiry int    `mapstructure:"REFRESH_TOKEN_EXPIRY"`

	// Mongo
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

func LoadEnvConfigs(env string) *envConfigs {
	cfg := new(envConfigs)

	if env == Development {
		cfg.LoadFromEnvFile()
	} else if env == Production {
		cfg.LoadConfigsFromOS()
	} else {
		log.Fatal().Msg("invalid environment")
	}

	return cfg
}

func (cfg *envConfigs) LoadFromEnvFile() {
	viper.AddConfigPath(".")
	viper.SetConfigFile("dev.env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal().Err(err).Msg("cannot load environment variables from envConfigs file")
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatal().Err(err).Msg("unmarshal envConfigs error happened")
	}
}

func (cfg *envConfigs) LoadConfigsFromOS() {
	t := reflect.ValueOf(&cfg).Elem().Type()
	for i := 0; i < t.NumField(); i++ {
		if tagValue := t.Field(i).Tag.Get("mapstructure"); tagValue != "-" {
			if err := viper.BindEnv(tagValue); err != nil {
				log.Fatal().Err(err).Msgf("cannot bind envConfigs: %v", tagValue)
			}
		}
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatal().Err(err).Msg("unmarshal envConfigs error happened")
	}
}
