package labelrepo

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type labelRepository struct {
	db *mongo.Database
}

func NewLabelRepository(db *mongo.Database) *labelRepository {
	return &labelRepository{db: db}
}
