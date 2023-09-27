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
}

type envConfigs struct {
	env struct {
		Port         int    `mapstructure:"PORT"`
		MongoConnStr string `mapstructure:"MONGO_CONN_STRING"`
		MongoDBName  string `mapstructure:"MONGO_DATABASE_NAME"`
		RedisAddr    string `mapstructure:"REDIS_ADDR"`
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
