package userrepo

import (
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
	"pro-magnet/common"
	usermodel "pro-magnet/modules/user/model"
)

func (repo *userRepository) Find(
	ctx context.Context,
	filter map[string]interface{},
	opts ...*options.FindOptions,
) ([]usermodel.User, error) {
	cursor, err := repo.db.Collection(usermodel.UserCollectionName).Find(ctx, filter, opts...)
	if err != nil {
		return nil, common.NewServerErr(err)
	}

	result := make([]usermodel.User, 0)
	if err = cursor.All(ctx, &result); err != nil {
		return nil, common.NewServerErr(err)
	}

	return result, nil
}

func (repo *userRepository) FindSimpleUsersByIds(
	ctx context.Context,
	userIds []string,
) ([]usermodel.User, error) {
	userOids := make([]primitive.ObjectID, 0)
	for _, id := range userIds {
		oid, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return nil, common.NewBadRequestErr(errors.New("invalid objectId"))
		}
		userOids = append(userOids, oid)
	}

	projectionOpt := options.Find().SetProjection(bson.M{
		"_id":    1,
		"name":   1,
		"email":  1,
		"avatar": 1,
	})

	return repo.Find(ctx, map[string]interface{}{"_id": bson.M{"$in": userOids}}, projectionOpt)
}

func (repo *userRepository) FindUsersByMatchingAtLeastOneCardSkills(
	ctx context.Context,
	boardId string, skills []string,
	exceptedUsedIds []string,
) ([]usermodel.User, error) {
	boardOid, err := primitive.ObjectIDFromHex(boardId)
	if err != nil {
		return nil, common.NewBadRequestErr(errors.New("invalid objectId"))
	}
	exceptedUserOids := make([]primitive.ObjectID, 0)
	for _, id := range exceptedUsedIds {
		oid, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return nil, common.NewBadRequestErr(errors.New("invalid objectId"))
		}
		exceptedUserOids = append(exceptedUserOids, oid)
	}
	aggPipeline := bson.A{
		bson.M{"$lookup": bson.M{
			"from":         "board_members",
			"foreignField": "userId",
			"localField":   "_id",
			"as":           "boardMembers",
		}},
		bson.M{"$unwind": "$boardMembers"},
		bson.M{"$match": bson.M{
			"_id":                  bson.M{"$nin": exceptedUserOids},
			"boardMembers.boardId": boardOid,
			"skills":               bson.M{"$in": skills}},
		},
		bson.M{"$project": bson.M{
			"_id":    1,
			"email":  1,
			"avatar": 1,
			"skills": 1,
		}},
	}

	cursor, err := repo.db.Collection(usermodel.UserCollectionName).Aggregate(ctx, aggPipeline)
	if err != nil {
		return nil, common.NewServerErr(err)
	}

	result := make([]usermodel.User, 0)
	if err = cursor.All(ctx, &result); err != nil {
		return nil, common.NewServerErr(err)
	}

	return result, nil
}
