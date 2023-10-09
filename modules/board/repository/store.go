package boardrepo

import "go.mongodb.org/mongo-driver/mongo"

type boardRepository struct {
	db *mongo.Database
}

func NewBoardRepository(db *mongo.Database) *boardRepository {
	return &boardRepository{db: db}
}
