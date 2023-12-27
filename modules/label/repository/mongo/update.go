package labelrepo

import (
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
	"pro-magnet/common"
	labelmodel "pro-magnet/modules/label/model"
)

func (repo *labelRepository) Update(
	ctx context.Context,
	filter map[string]interface{},
	updateData *labelmodel.LabelUpdate,
) error {
	result, err := repo.db.
		Collection(labelmodel.LabelCollectionName).
		UpdateOne(ctx, filter, bson.M{
			"$set": updateData,
		})
	if err != nil {
		return common.NewServerErr(err)
	}
	if result.MatchedCount == 0 {
		return common.NewNotFoundErr("label", mongo.ErrNoDocuments)
	}

	return nil
}

func (repo *labelRepository) UpdateById(
	ctx context.Context,
	labelId string,
	updateData *labelmodel.LabelUpdate,
) error {
	oid, err := primitive.ObjectIDFromHex(labelId)
	if err != nil {
		return common.NewBadRequestErr(errors.New("invalid object Id"))
	}

	return repo.Update(ctx, map[string]interface{}{"_id": oid}, updateData)
}
