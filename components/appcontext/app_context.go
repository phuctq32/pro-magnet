package appcontext

import (
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"pro-magnet/components/validator"
)

type AppContext interface {
	DBConnection() *mongo.Database
	RedisClient() *redis.Client
	Validator() validator.Validator
}

type appCxt struct {
	mongodb   *mongo.Database
	redisCli  *redis.Client
	validator validator.Validator
}

func NewAppContext(
	db *mongo.Database,
	redisCli *redis.Client,
	validator validator.Validator,
) AppContext {
	return &appCxt{
		mongodb:   db,
		redisCli:  redisCli,
		validator: validator,
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
