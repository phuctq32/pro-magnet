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
	VerificationLink() string
}

type envConfigs struct {
	env struct {
		// app
		Port     int    `mapstructure:"PORT"`
		FeDomain string `mapstructure:"FE_DOMAIN"`

		// mongodb
		MongoConnStr string `mapstructure:"MONGO_CONN_STRING"`
		MongoDBName  string `mapstructure:"MONGO_DATABASE_NAME"`

		// redis
		RedisAddr string `mapstructure:"REDIS_ADDR"`

		// sendgrid
		SendGridApiKey                string `mapstructure:"SENDGRID_API_KEY"`
		SendGridFromEmail             string `mapstructure:"FROM_EMAIL"`
		SendGridVerifyEmailTemplateId string `mapstructure:"SENDGRID_VERIFY_TEMPlATE_ID"`
		VerificationLink              string `mapstructure:"VERIFICATION_LINK"`
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

func (cfg *envConfigs) VerificationLink() string {
	return cfg.env.VerificationLink
}
