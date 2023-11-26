package columnrepo

import (
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/net/context"
	"pro-magnet/common"
	columnmodel "pro-magnet/modules/column/model"
)

// Delete do the soft delete operation
func (repo *columnRepository) Delete(
	ctx context.Context,
	filter map[string]interface{},
) error {
	_, err := repo.db.
		Collection(columnmodel.ColumnCollectionName).
		UpdateMany(ctx, filter, bson.M{
			"$set": bson.M{
				"status": columnmodel.Deleted,
			},
		})
	if err != nil {
		return common.NewServerErr(err)
	}

	return nil
}

func (repo *columnRepository) DeleteById(
	ctx context.Context,
	id string,
) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return common.NewBadRequestErr(errors.New("invalid objectId"))
	}

	return repo.Delete(ctx, map[string]interface{}{"_id": oid})
}
