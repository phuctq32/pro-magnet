package wsmemberrepo

import "go.mongodb.org/mongo-driver/mongo"

type wsMemberRepository struct {
	db *mongo.Database
}

func NewWorkspaceMemberRepository(db *mongo.Database) *wsMemberRepository {
	return &wsMemberRepository{db: db}
}
