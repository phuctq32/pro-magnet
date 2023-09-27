package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"os"
	"pro-magnet/components/appcontext"
	"pro-magnet/configs"
	"pro-magnet/middlewares"
	"time"
)

func main() {
	env := os.Getenv("ENV")

	initLogger(env)

	// Load env
	configs.LoadEnvConfigs(env)
	log.Info().Interface("env", configs.EnvConfigs.MongoConnectionString()).Msg("")

	// DB
	db, cancel := connectMongoDB()
	defer cancel()

	// Redis Client
	redisCli := connectRedisCli()

	// Init AppContext
	_ = appcontext.NewAppContext(db, redisCli)

	if env == configs.Production {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()

	if env == configs.Development {
		router.Use(gin.Recovery())
	}

	router.Use(middlewares.Logger(), middlewares.Recover())

	router.GET("/ping", func(c *gin.Context) {
		a := []int{1}
		_ = a[2]

		log.Debug().Err(errors.New("hello")).Msg("")
		c.JSON(http.StatusOK, gin.H{"message": "OK"})
	})

	err := router.Run(fmt.Sprintf(":%v", configs.EnvConfigs.Port()))
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
}

func connectMongoDB() (*mongo.Database, context.CancelFunc) {
	mongoUri := configs.EnvConfigs.MongoConnectionString()
	mongoDBName := configs.EnvConfigs.MongoDBName()

	opts := options.Client().ApplyURI(mongoUri)
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*10)
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		log.Fatal().Msg(err.Error())
	}

	// Check connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal().Err(err).Msg("connect to mongodb failed")
	}

	db := client.Database(mongoDBName)
	log.Debug().Msg("db connected")

	return db, cancel
}

func initLogger(env string) {
	zerolog.TimeFieldFormat = time.RFC3339
	zerolog.SetGlobalLevel(zerolog.InfoLevel) // production

	if env == configs.Development {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	// Add file and line number to log
	log.Logger = log.With().Caller().Logger()
}

func connectRedisCli() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     configs.EnvConfigs.RedisAddr(),
		Password: "", // no password set
		DB:       0,  // default db
	})
}
