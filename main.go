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
	"pro-magnet/configs"
	"pro-magnet/middlewares"
	"pro-magnet/routes"
	"time"
)

func main() {
	env := os.Getenv("ENV")

	initLogger(env)

	// Load env
	configs.LoadEnvConfigs(env)

	// DB
	db, cancel := connectMongoDB()
	defer cancel()

	// Redis Client
	redisCli := connectRedisCli()

	// Validator
	appValidator := validator.NewValidator()

	// S3 upload provider
	s3Uploader := upload.NewS3Provider(
		configs.EnvConfigs.S3AccessKey(),
		configs.EnvConfigs.S3SecretKey(),
		configs.EnvConfigs.S3BucketName(),
		configs.EnvConfigs.S3Region(),
		configs.EnvConfigs.S3Domain(),
	)

	// Cloudinary uploade provider
	cldUploader, err := upload.NewCloudinaryUploader(
		configs.EnvConfigs.CloudinaryCloudName(),
		configs.EnvConfigs.CloudinaryApiKey(),
		configs.EnvConfigs.CloudinaryApiSecret(),
	)
	if err != nil {
		log.Fatal().Err(err)
	}

	// Async Group
	asyncg, agCancel := asyncgroup.New(10000)
	defer agCancel()

	// Init AppContext
	appCtx := appcontext.NewAppContext(db, redisCli, appValidator, asyncg, s3Uploader, cldUploader)

	if env == configs.Production {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()

	router.Use(middlewares.Logger())
	if env == configs.Development {
		router.Use(gin.Recovery())
	}

	router.Use(middlewares.Recover())

	router.Use(Cors("*"))

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "OK"})
	})

	// setup routes
	routes.Setup(appCtx, router)

	err = router.Run(fmt.Sprintf(":%v", configs.EnvConfigs.Port()))
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
