package wsrepo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"pro-magnet/common"
	wsmodel "pro-magnet/modules/workspace/model"
	"time"
)

func (repo *wsRepository) Create(
	ctx context.Context,
	data *wsmodel.WorkspaceCreation,
) (*wsmodel.Workspace, error) {
	ownerOid, _ := primitive.ObjectIDFromHex(data.OwnerUserId)
	now := time.Now()

	wsInsert := &wsmodel.WorkspaceInsert{
		CreatedAt:   now,
		UpdatedAt:   now,
		OwnerUserId: ownerOid,
		Name:        data.Name,
		Image:       data.Image,
	}

	result, err := repo.db.
		Collection(wsmodel.WsCollectionName).
		InsertOne(ctx, wsInsert)
	if err != nil {
		return nil, common.NewServerErr(err)
	}

	insertedId := result.InsertedID.(primitive.ObjectID).Hex()

	return &wsmodel.Workspace{
		Id:          &insertedId,
		CreatedAt:   wsInsert.CreatedAt,
		UpdatedAt:   wsInsert.UpdatedAt,
		Name:        wsInsert.Name,
		Image:       wsInsert.Image,
		OwnerUserId: data.OwnerUserId,
	}, nil
}
