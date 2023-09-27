package appcontext

import (
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
)

type AppContext interface {
	DBConnection() *mongo.Database
	RedisClient() *redis.Client
}

type appCxt struct {
	mongodb  *mongo.Database
	redisCli *redis.Client
}

func NewAppContext(
	db *mongo.Database,
	redisCli *redis.Client,
) AppContext {
	return &appCxt{
		mongodb:  db,
		redisCli: redisCli,
	}
}

func (appCtx *appCxt) DBConnection() *mongo.Database {
	return appCtx.mongodb
}

func (appCtx *appCxt) RedisClient() *redis.Client {
	return appCtx.redisCli
}
