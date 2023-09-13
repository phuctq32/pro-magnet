package configs

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"os"
	"reflect"
	"strconv"
)

const (
	Development string = "development"
	Production         = "production"
)

type EnvConfiguration interface {
	GetPort() int
	GetMongoConnectionString() string
	GetMongoDBName() string
}

var EnvConfigs EnvConfiguration

type envConfigs struct {
	// app
	Port int `mapstructure:"PORT"`

	// mongodb
	MongoConnStr string `mapstructure:"MONGO_CONN_STRING"`
	MongoDBName  string `mapstructure:"MONGO_DATABASE_NAME"`
}

func LoadEnvConfigs(env string) {
	if env == Development {
		EnvConfigs = LoadConfigsFromEnvFile()
	} else if env == Production {
		EnvConfigs = LoadConfigsFromOS()
	} else {
		log.Fatal().Msg("invalid environment")
	}
}

func LoadConfigsFromEnvFile() EnvConfiguration {
	var cfg envConfigs
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal().Err(err).Msg("cannot load environment variables from env file")
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatal().Err(err).Msg("something went wrong while loading env")
	}

	return &cfg
}

func LoadConfigsFromOS() EnvConfiguration {
	var cfg envConfigs
	v := reflect.ValueOf(&cfg).Elem()
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		if tagValue := t.Field(i).Tag.Get("mapstructure"); tagValue != "-" {
			val := os.Getenv(tagValue)
			if val == "" {
				log.Fatal().Msgf("cannot get %v from env file", tagValue)
			}

			switch v.Field(i).Kind() {
			case reflect.Int:
				intVal, err := strconv.Atoi(val)
				if err != nil {
					log.Fatal().Msgf("cannot convert string to int %v", tagValue)
				}
				v.Field(i).SetInt(int64(intVal))
			default:
				v.Field(i).Set(reflect.ValueOf(val))
			}
		}
	}

	return &cfg
}

func (cfg *envConfigs) GetPort() int {
	return cfg.Port
}

func (cfg *envConfigs) GetMongoConnectionString() string {
	return cfg.MongoConnStr
}

func (cfg *envConfigs) GetMongoDBName() string {
	return cfg.MongoDBName
}
