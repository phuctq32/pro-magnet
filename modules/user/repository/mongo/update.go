package userrepo

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"pro-magnet/common"
	usermodel "pro-magnet/modules/user/model"
)

func (repo *userRepository) Update(
	ctx context.Context,
	filter map[string]interface{},
	updateData map[string]interface{},
) (*usermodel.User, error) {
	var updatedUser usermodel.User

	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	if err := repo.db.
		Collection(usermodel.UserCollectionName).
		FindOneAndUpdate(ctx, filter, bson.M{
			"$set": updateData,
			"$currentDate": bson.M{
				"updatedAt": bson.M{"$type": "date"},
			},
		}, opts).Decode(&updatedUser); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, common.NewNotFoundErr("user", err)
		}
		return nil, common.NewServerErr(err)
	}

	return &updatedUser, nil
}

func (repo *userRepository) UpdateById(
	ctx context.Context,
	id string,
	updateData map[string]interface{},
) (*usermodel.User, error) {
	oid, _ := primitive.ObjectIDFromHex(id)
	return repo.Update(ctx, bson.M{"_id": oid}, updateData)
}

func (repo *userRepository) SetEmailVerified(
	ctx context.Context,
	id string,
) error {
	if _, err := repo.UpdateById(ctx, id, map[string]interface{}{"isVerified": true}); err != nil {
		return err
	}

	return nil
}

func (repo *userRepository) UpdatePasswordByEmail(
	ctx context.Context,
	email string,
	password string,
) error {
	if _, err := repo.Update(
		ctx,
		map[string]interface{}{"email": email},
		map[string]interface{}{"password": password},
	); err != nil {
		return err
	}

	return nil
}
