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
			"$set": updateData.ToUpdateData(),
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

func (repo *columnRepository) RemoveCardId(
	ctx context.Context,
	columnId, cardId string,
) error {
	columnOid, err := primitive.ObjectIDFromHex(columnId)
	if err != nil {
		return common.NewBadRequestErr(errors.New("invalid objectId"))
	}
	cardOid, err := primitive.ObjectIDFromHex(cardId)
	if err != nil {
		return common.NewBadRequestErr(errors.New("invalid objectId"))
	}

	result, err := repo.db.
		Collection(columnmodel.ColumnCollectionName).
		UpdateOne(ctx, bson.M{"_id": columnOid},
			bson.M{
				"$pull": bson.M{"orderedCardIds": cardOid},
			},
		)
	if err != nil {
		return common.NewServerErr(err)
	}
	if result.MatchedCount == 0 {
		return common.NewNotFoundErr("column", mongo.ErrNoDocuments)
	}

	return nil
}
