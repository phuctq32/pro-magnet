package wsrepo

import (
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
	"pro-magnet/common"
	wsmodel "pro-magnet/modules/workspace/model"
	"time"
)

func (repo *wsRepository) Update(
	ctx context.Context,
	filter map[string]interface{},
	updateData *wsmodel.WorkspaceUpdate,
) error {
	updateData.UpdatedAt = time.Now()

	result, err := repo.db.Collection(wsmodel.WsCollectionName).UpdateOne(ctx, filter, bson.M{
		"$set": updateData,
	})
	if err != nil {
		return common.NewServerErr(err)
	}
	if result.MatchedCount == 0 {
		return common.NewNotFoundErr("workspace", mongo.ErrNoDocuments)
	}

	return nil
}

func (repo *wsRepository) UpdateById(
	ctx context.Context,
	workspaceId string,
	updateData *wsmodel.WorkspaceUpdate,
) error {
	oid, err := primitive.ObjectIDFromHex(workspaceId)
	if err != nil {
		return common.NewBadRequestErr(errors.New("invalid object Id"))
	}

	return repo.Update(ctx, map[string]interface{}{"_id": oid}, updateData)
}
