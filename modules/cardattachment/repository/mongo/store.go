package carepo

import "go.mongodb.org/mongo-driver/mongo"

type cardAttachmentRepository struct {
	db *mongo.Database
}

func NewCardAttachmentRepository(db *mongo.Database) *cardAttachmentRepository {
	return &cardAttachmentRepository{db: db}
}
