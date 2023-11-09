package wsrepo

import "go.mongodb.org/mongo-driver/mongo"

type wsRepository struct {
	db *mongo.Database
}

func NewWorkspaceRepository(db *mongo.Database) *wsRepository {
	return &wsRepository{db: db}
}
