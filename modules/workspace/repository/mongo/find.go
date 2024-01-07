package wsrepo

import (
	"context"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"pro-magnet/common"
	wsmodel "pro-magnet/modules/workspace/model"
)

func (repo *wsRepository) Find(
	ctx context.Context,
	filter map[string]interface{},
) ([]wsmodel.Workspace, error) {
	cursor, err := repo.db.
		Collection(wsmodel.WsCollectionName).
		Find(ctx, filter)
	if err != nil {
		return nil, common.NewServerErr(err)
	}

	result := make([]wsmodel.Workspace, 0)
	if err = cursor.All(ctx, &result); err != nil {
		return nil, common.NewServerErr(err)
	}

	return result, nil
}

func (repo *wsRepository) FindByIds(
	ctx context.Context,
	wsIds []string,
) ([]wsmodel.Workspace, error) {
	wsOids := make([]primitive.ObjectID, 0)
	for _, id := range wsIds {
		oid, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return nil, common.NewBadRequestErr(errors.New("invalid objectId"))
		}
		wsOids = append(wsOids, oid)
	}

	return repo.Find(ctx, map[string]interface{}{"_id": bson.M{"$in": wsOids}})
}

func (repo *wsRepository) FindOne(
	ctx context.Context,
	filter map[string]interface{},
) (*wsmodel.Workspace, error) {
	var result wsmodel.Workspace
	if err := repo.db.
		Collection(wsmodel.WsCollectionName).
		FindOne(ctx, filter).
		Decode(&result); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, common.NewNotFoundErr("workspace", err)
		}

		return nil, common.NewServerErr(err)
	}

	return &result, nil
}

func (repo *wsRepository) FindById(
	ctx context.Context,
	id string,
) (*wsmodel.Workspace, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, common.NewBadRequestErr(errors.New("invalid object Id"))
	}

	return repo.FindOne(ctx, map[string]interface{}{"_id": oid})
}

func (repo *wsRepository) FindByName(
	ctx context.Context,
	name string,
) (*wsmodel.Workspace, error) {
	return repo.FindOne(ctx, map[string]interface{}{"name": name})
}

func (repo *wsRepository) Search(
	ctx context.Context,
	memberId, searchTerm string,
) ([]wsmodel.Workspace, error) {
	memberOid, err := primitive.ObjectIDFromHex(memberId)
	if err != nil {
		return nil, common.NewBadRequestErr(errors.New("invalid objectId"))
	}

	aggPipeline := bson.A{
		bson.M{"$lookup": bson.M{
			"from":         "workspace_members",
			"foreignField": "workspaceId",
			"localField":   "_id",
			"as":           "wsMembers",
		}},
		bson.M{"$unwind": "$wsMembers"},
		bson.M{"$match": bson.M{
			"wsMembers.userId": memberOid,
			"name": primitive.Regex{
				Pattern: searchTerm,
				Options: "i",
			},
		}},
	}

	cursor, err := repo.db.Collection(wsmodel.WsCollectionName).Aggregate(ctx, aggPipeline)
	if err != nil {
		return nil, common.NewServerErr(err)
	}

	result := make([]wsmodel.Workspace, 0)
	if err = cursor.All(ctx, &result); err != nil {
		return nil, common.NewServerErr(err)
	}

	return result, nil
}
