package appcontext

import (
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"pro-magnet/components/upload"
	"pro-magnet/components/validator"
)

type AppContext interface {
	DBConnection() *mongo.Database
	RedisClient() *redis.Client
	Validator() validator.Validator
	S3Uploader() upload.Uploader
}

type appCxt struct {
	mongodb    *mongo.Database
	redisCli   *redis.Client
	validator  validator.Validator
	s3Uploader upload.Uploader
}

func NewAppContext(
	db *mongo.Database,
	redisCli *redis.Client,
	validator validator.Validator,
	s3Uploader upload.Uploader,
) AppContext {
	return &appCxt{
		mongodb:    db,
		redisCli:   redisCli,
		validator:  validator,
		s3Uploader: s3Uploader,
	}
}

func (appCtx *appCxt) DBConnection() *mongo.Database {
	return appCtx.mongodb
}

func (appCtx *appCxt) RedisClient() *redis.Client {
	return appCtx.redisCli
}

func (appCtx *appCxt) Validator() validator.Validator {
	return appCtx.validator
}

func (appCtx *appCxt) S3Uploader() upload.Uploader {
	return appCtx.s3Uploader
}
