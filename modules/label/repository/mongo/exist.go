package labelrepo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
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
	labelId *string,
	boardId, title, color string,
) (bool, error) {
	boardOid, err := primitive.ObjectIDFromHex(boardId)
	if err != nil {
		return false, common.NewBadRequestErr(err)
	}
	filter := map[string]interface{}{
		"status":  labelmodel.Active,
		"boardId": boardOid,
		"title":   title,
		"color":   color,
	}
	if labelId != nil {
		labelOid, err := primitive.ObjectIDFromHex(*labelId)
		if err != nil {
			return false, common.NewBadRequestErr(err)
		}
		filter["_id"] = bson.M{"$ne": labelOid}
	}

	return repo.Exist(ctx, filter)
}
