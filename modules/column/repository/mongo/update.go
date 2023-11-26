package columnrepo

import (
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
	"pro-magnet/common"
	columnmodel "pro-magnet/modules/column/model"
)

func (repo *columnRepository) Update(
	ctx context.Context,
	filter map[string]interface{},
	updateData *columnmodel.ColumnUpdate,
) (*columnmodel.Column, error) {
	var updatedCol columnmodel.Column

	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	if err := repo.db.
		Collection(columnmodel.ColumnCollectionName).
		FindOneAndUpdate(ctx, filter, bson.M{
			"$set": updateData,
			"$currentDate": bson.M{
				"updatedAt": bson.M{"$type": "date"},
			},
		}, opts).Decode(&updatedCol); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, common.NewNotFoundErr("column", err)
		}
		return nil, common.NewServerErr(err)
	}

	return &updatedCol, nil
}

func (repo *columnRepository) UpdateById(
	ctx context.Context,
	id string,
	updateData *columnmodel.ColumnUpdate,
) (*columnmodel.Column, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, common.NewBadRequestErr(errors.New("invalid objectId"))
	}

	return repo.Update(ctx, map[string]interface{}{"_id": oid}, updateData)
}
