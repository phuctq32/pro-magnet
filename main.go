package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
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

	// load env
	configs.LoadEnvConfigs(env)

	// db
	db, cancel := connectMongoDB()
	defer cancel()

	// init AppContext
	_ = appcontext.NewAppContext(db)

	if env == configs.Production {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()

	router.Use(middlewares.Logger(), middlewares.Recover())

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "OK"})
	})

	err := router.Run(fmt.Sprintf(":%v", configs.EnvConfigs.GetPort()))
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
}

func connectMongoDB() (*mongo.Database, context.CancelFunc) {
	mongoUri := configs.EnvConfigs.GetMongoConnectionString()
	mongoDBName := configs.EnvConfigs.GetMongoDBName()

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

	// add file and line number to log
	log.Logger = log.With().Caller().Logger()
}
