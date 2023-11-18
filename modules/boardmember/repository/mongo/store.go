package boardmemberrepo

import "go.mongodb.org/mongo-driver/mongo"

type boardMemberRepository struct {
	db *mongo.Database
}

func NewBoardMemberRepository(db *mongo.Database) *boardMemberRepository {
	return &boardMemberRepository{db: db}
}
