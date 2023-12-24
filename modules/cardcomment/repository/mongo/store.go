package cardcommentrepo

import "go.mongodb.org/mongo-driver/mongo"

type cardCommentRepository struct {
	db *mongo.Database
}

func NewCardCommentRepository(db *mongo.Database) *cardCommentRepository {
	return &cardCommentRepository{db: db}
}
