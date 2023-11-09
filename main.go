package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"os"
	"pro-magnet/components/appcontext"
	"pro-magnet/components/asyncgroup"
	"pro-magnet/components/upload"
	"pro-magnet/components/validator"
	"pro-magnet/configs/envconfigs"
	"pro-magnet/middlewares"
	"pro-magnet/routes"
	"time"
)

func main() {
	env := os.Getenv("ENV")
	initLogger(env)

	// Load env configs
	envConfigs := envconfigs.New(env)

	// DB
	db, cancel := connectMongoDB(envConfigs.Mongo().ConnectionString(), envConfigs.Mongo().DBName())
	defer cancel()

	// Redis Client
	redisCli := connectRedisCli(envConfigs.Redis().Address())

	// Validator
	appValidator := validator.NewValidator()

	// S3 upload provider
	s3Uploader := upload.NewS3Provider(
		envConfigs.AwsS3().AccessKey(),
		envConfigs.AwsS3().SecretKey(),
		envConfigs.AwsS3().BucketName(),
		envConfigs.AwsS3().Region(),
		envConfigs.AwsS3().Domain(),
	)

	// Cloudinary upload provider
	cldUploader, err := upload.NewCloudinaryUploader(
		envConfigs.Cloudinary().CloudName(),
		envConfigs.Cloudinary().ApiKey(),
		envConfigs.Cloudinary().ApiSecret(),
	)
	if err != nil {
		log.Fatal().Err(err)
	}

	// Async Group
	asyncg, agCancel := asyncgroup.New(10000)
	defer agCancel()

	// Init AppContext
	appCtx := appcontext.NewAppContext(envConfigs, db, redisCli, appValidator, asyncg, s3Uploader, cldUploader)

	if env == envconfigs.Production {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()

	router.Use(middlewares.Logger())
	if env == envconfigs.Development {
		router.Use(gin.Recovery())
	}

	router.Use(middlewares.Recover())

	router.Use(Cors("*"))

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "OK"})
	})

	// setup routes
	routes.Setup(appCtx, router)

	err = router.Run(fmt.Sprintf(":%v", envConfigs.App().Port()))
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
}

func Cors(allowOrigin string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", allowOrigin)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, Content-Length, X-CSRF-Token, Token, session, Origin, Host, Connection, Accept-Encoding, Accept-Language, X-Requested-With")

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Request.Header.Del("Origin")

		c.Next()
	}
}

func connectMongoDB(mongoUri, mongoDBName string) (*mongo.Database, context.CancelFunc) {
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

	if env == envconfigs.Development {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	// Add file and line number to log
	log.Logger = log.With().Caller().Logger()
}

func connectRedisCli(redisAddr string) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: "", // no password set
		DB:       0,  // default db
	})
}
