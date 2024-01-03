package boardrepo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"pro-magnet/common"
	boardmodel "pro-magnet/modules/board/model"
)

func (repo *boardRepository) Exists(
	ctx context.Context,
	filter map[string]interface{},
) (bool, error) {
	count, err := repo.db.
		Collection(boardmodel.BoardCollectionName).
		CountDocuments(ctx, filter)
	if err != nil {
		return false, common.NewServerErr(err)
	}

	return count > 0, nil
}

func (repo *boardRepository) ExistsInWorkspace(
	ctx context.Context,
	boardId *string,
	boardName string,
	workspaceId string,
) (bool, error) {
	wsOid, _ := primitive.ObjectIDFromHex(workspaceId)
	filter := map[string]interface{}{"workspaceId": wsOid, "name": boardName}
	if boardId != nil {
		boardOid, _ := primitive.ObjectIDFromHex(*boardId)
		filter["_id"] = bson.M{"$ne": boardOid}
	}
	return repo.Exists(ctx, filter)
}
