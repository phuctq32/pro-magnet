package mongo

import "go.mongodb.org/mongo-driver/mongo"

type cardRepository struct {
	db *mongo.Database
}

func NewCardRepository(db *mongo.Database) *cardRepository {
	return &cardRepository{db: db}
}
