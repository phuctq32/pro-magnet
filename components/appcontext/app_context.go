package appcontext

import (
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"pro-magnet/components/asyncgroup"
	"pro-magnet/components/upload"
	"pro-magnet/components/validator"
	"pro-magnet/configs/envconfigs"
)

type AppContext interface {
	EnvConfigs() envconfigs.Configs
	DBConnection() *mongo.Database
	RedisClient() *redis.Client
	Validator() validator.Validator
	S3Uploader() upload.Uploader
	AsyncGroup() asyncgroup.AsyncGroup
	CloudinaryUploader() upload.Uploader
}

type appContext struct {
	envConfigs  envconfigs.Configs
	mongodb     *mongo.Database
	redisCli    *redis.Client
	validator   validator.Validator
	asyncg      asyncgroup.AsyncGroup
	s3Uploader  upload.Uploader
	cldUploader upload.Uploader
}

func NewAppContext(
	envConfigs envconfigs.Configs,
	db *mongo.Database,
	redisCli *redis.Client,
	validator validator.Validator,
	asyncg asyncgroup.AsyncGroup,
	s3Uploader upload.Uploader,
	cldUploader upload.Uploader,
) AppContext {
	return &appContext{
		envConfigs:  envConfigs,
		mongodb:     db,
		redisCli:    redisCli,
		validator:   validator,
		asyncg:      asyncg,
		s3Uploader:  s3Uploader,
		cldUploader: cldUploader,
	}
}

func (appCtx *appContext) EnvConfigs() envconfigs.Configs {
	return appCtx.envConfigs
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

func (appCtx *appContext) CloudinaryUploader() upload.Uploader {
	return appCtx.cldUploader
}
