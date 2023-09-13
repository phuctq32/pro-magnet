package appcontext

import "go.mongodb.org/mongo-driver/mongo"

type AppContext interface {
	GetDBConn() *mongo.Database
}

type appCxt struct {
	mongodb *mongo.Database
}

func NewAppContext(db *mongo.Database) AppContext {
	return &appCxt{mongodb: db}
}

func (appCtx *appCxt) GetDBConn() *mongo.Database {
	return appCtx.mongodb
}
