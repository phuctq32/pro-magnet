package appcontext

import (
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"pro-magnet/components/asyncgroup"
	"pro-magnet/components/validator"
)

type AppContext interface {
	DBConnection() *mongo.Database
	RedisClient() *redis.Client
	Validator() validator.Validator
	AsyncGroup() asyncgroup.AsyncGroup
}

type appContext struct {
	mongodb   *mongo.Database
	redisCli  *redis.Client
	validator validator.Validator
	asyncg    asyncgroup.AsyncGroup
}

func NewAppContext(
	db *mongo.Database,
	redisCli *redis.Client,
	validator validator.Validator,
	asyncg asyncgroup.AsyncGroup,
) AppContext {
	return &appContext{
		mongodb:   db,
		redisCli:  redisCli,
		validator: validator,
		asyncg:    asyncg,
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
