package cardchecklistrepo

import "go.mongodb.org/mongo-driver/mongo"

type cardChecklistRepository struct {
	db *mongo.Database
}

func NewCardChecklistRepository(
	db *mongo.Database,
) *cardChecklistRepository {
	return &cardChecklistRepository{db: db}
}
