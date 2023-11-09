package wsrepo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"pro-magnet/common"
	wsmodel "pro-magnet/modules/workspace/model"
)

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

func (repo *wsRepository) FindByName(
	ctx context.Context,
	name string,
) (*wsmodel.Workspace, error) {
	return repo.FindOne(ctx, map[string]interface{}{"name": name})
}
