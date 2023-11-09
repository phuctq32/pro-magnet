package labelrepo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"pro-magnet/common"
	labelmodel "pro-magnet/modules/label/model"
)

func (repo *labelRepository) Exist(
	ctx context.Context,
	filter map[string]interface{},
) (bool, error) {
	count, err := repo.db.
		Collection(labelmodel.LabelCollectionName).
		CountDocuments(ctx, filter)
	if err != nil {
		return false, common.NewServerErr(err)
	}

	return count > 0, nil
}

func (repo *labelRepository) ExistsInBoard(
	ctx context.Context,
	data *labelmodel.LabelCreation,
) (bool, error) {
	boardOid, err := primitive.ObjectIDFromHex(data.BoardId)
	if err != nil {
		return false, common.NewBadRequestErr(err)
	}

	return repo.Exist(ctx, map[string]interface{}{
		"boardId": boardOid,
		"title":   data.Title,
		"color":   data.Color,
	})
}
