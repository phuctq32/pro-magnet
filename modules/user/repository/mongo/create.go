package userrepo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"pro-magnet/common"
	usermodel "pro-magnet/modules/user/model"
)

func (repo *userRepository) Create(ctx context.Context, data *usermodel.User) (*string, error) {
	result, err := repo.db.Collection(usermodel.UserCollectionName).InsertOne(ctx, data)
	if err != nil {
		return nil, common.NewServerErr(err)
	}
	id := result.InsertedID.(primitive.ObjectID).Hex()

	return &id, nil
}
