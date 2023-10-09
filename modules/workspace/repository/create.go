package wsrepo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"pro-magnet/common"
	wsmodel "pro-magnet/modules/workspace/model"
)

func (repo *wsRepository) Create(
	ctx context.Context,
	data *wsmodel.Workspace,
) (*wsmodel.Workspace, error) {
	userOid, _ := primitive.ObjectIDFromHex(data.UserId)

	wsInsert := &wsmodel.WorkspaceInsert{
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
		UserId:    userOid,
		Name:      data.Name,
		Image:     data.Image,
	}

	result, err := repo.db.
		Collection(wsmodel.WsCollectionName).
		InsertOne(ctx, wsInsert)
	if err != nil {
		return nil, common.NewServerErr(err)
	}
	insertedId := result.InsertedID.(primitive.ObjectID).Hex()
	data.Id = &insertedId

	return data, nil
}
