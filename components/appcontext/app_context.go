package appcontext

import (
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"pro-magnet/components/asyncgroup"
	"pro-magnet/components/upload"
	"pro-magnet/components/validator"
)

type AppContext interface {
	DBConnection() *mongo.Database
	RedisClient() *redis.Client
	Validator() validator.Validator
	S3Uploader() upload.Uploader
	AsyncGroup() asyncgroup.AsyncGroup
}

type appContext struct {
	mongodb    *mongo.Database
	redisCli   *redis.Client
	validator  validator.Validator
	s3Uploader upload.Uploader
	asyncg     asyncgroup.AsyncGroup
}

func NewAppContext(
	db *mongo.Database,
	redisCli *redis.Client,
	validator validator.Validator,
	asyncg asyncgroup.AsyncGroup,
	s3Uploader upload.Uploader,
) AppContext {
	return &appContext{
		mongodb:    db,
		redisCli:   redisCli,
		validator:  validator,
		asyncg:     asyncg,
		s3Uploader: s3Uploader,
	}
}

func (appCtx *appContext) DBConnection() *mongo.Database {
	return appCtx.mongodb
}

func (appCtx *appContext) RedisClient() *redis.Client {
	return appCtx.redisCli
}

func (appCtx *appContext) Validator() validator.Validator {
	return appCtx.validator
}

func (appCtx *appContext) AsyncGroup() asyncgroup.AsyncGroup {
	return appCtx.asyncg
}

func (appCtx *appContext) S3Uploader() upload.Uploader {
	return appCtx.s3Uploader
}
