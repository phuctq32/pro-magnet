package columnrepo

import "go.mongodb.org/mongo-driver/mongo"

type columnRepository struct {
	db *mongo.Database
}

func NewColumnRepository(db *mongo.Database) *columnRepository {
	return &columnRepository{db: db}
}
