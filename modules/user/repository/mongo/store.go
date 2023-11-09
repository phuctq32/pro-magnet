package userrepo

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct {
	db *mongo.Database
}

func NewUserRepository(db *mongo.Database) *userRepository {
	return &userRepository{db: db}
}
