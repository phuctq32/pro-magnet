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

func (repo *userRepository) UpdateSkills(
	ctx context.Context,
	userId string, skills []string,
) error {
	userOid, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return common.NewBadRequestErr(errors.New("invalid objectId"))
	}

	_, err = repo.db.Collection(usermodel.UserCollectionName).
		UpdateOne(
			ctx,
			bson.M{"_id": userOid},
			bson.M{"$set": bson.M{"skills": skills}},
		)

	return err
}

func (repo *userRepository) UpdateById(
	ctx context.Context,
	id string,
	updateData *usermodel.UserUpdate,
) (*usermodel.User, error) {
	oid, _ := primitive.ObjectIDFromHex(id)
	return updateUser(ctx, repo.db, bson.M{"_id": oid}, updateData)
}

func (repo *userRepository) SetEmailVerified(
	ctx context.Context,
	id string,
) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return common.NewBadRequestErr(errors.New("invalid objectId"))
	}
	if _, err = updateUser(
		ctx, repo.db,
		map[string]interface{}{"_id": oid},
		map[string]interface{}{"isVerified": true}); err != nil {
		return err
	}

	return nil
}

func (repo *userRepository) UpdatePasswordByEmail(
	ctx context.Context,
	email string,
	password string,
) error {
	if _, err := updateUser(
		ctx,
		repo.db,
		map[string]interface{}{"email": email},
		map[string]interface{}{"password": password},
	); err != nil {
		return err
	}

	return nil
}

func (repo *userRepository) UpdatePasswordById(
	ctx context.Context,
	id string,
	password string,
) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return common.NewBadRequestErr(errors.New("invalid objectId"))
	}
	if _, err := updateUser(
		ctx, repo.db,
		map[string]interface{}{"_id": oid},
		map[string]interface{}{"password": password}); err != nil {
		return err
	}

	return nil
}

func updateUser[T *usermodel.UserUpdate | map[string]interface{}](
	ctx context.Context,
	db *mongo.Database,
	filter map[string]interface{},
	updateData T,
) (*usermodel.User, error) {
	var updatedUser usermodel.User

	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	if err := db.
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
