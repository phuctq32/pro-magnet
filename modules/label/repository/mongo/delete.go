package labelrepo

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/net/context"
	"pro-magnet/common"
	labelmodel "pro-magnet/modules/label/model"
)

func (repo *labelRepository) DeleteOne(
	ctx context.Context,
	filter map[string]interface{},
) error {
	if _, err := repo.db.
		Collection(labelmodel.LabelCollectionName).
		UpdateOne(ctx, filter, bson.M{
			"$set": bson.M{"status": labelmodel.Deleted},
		}); err != nil {
		return common.NewServerErr(err)
	}

	return nil
}

func (repo *labelRepository) DeleteById(
	ctx context.Context,
	labelId string,
) error {
	labelOid, err := primitive.ObjectIDFromHex(labelId)
	if err != nil {
		return common.NewBadRequestErr(err)
	}

	return repo.DeleteOne(ctx, map[string]interface{}{"_id": labelOid})
}
