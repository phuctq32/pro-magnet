package carepo

import (
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/net/context"
	"pro-magnet/common"
	camodel "pro-magnet/modules/cardattachment/model"
)

// Delete do the soft delete operation
func (repo *cardAttachmentRepository) Delete(
	ctx context.Context,
	filter map[string]interface{},
) error {
	_, err := repo.db.
		Collection(camodel.CardAttachmentCollectionName).
		UpdateMany(ctx, filter, bson.M{
			"$set": bson.M{
				"status": camodel.Deleted,
			},
		})
	if err != nil {
		return common.NewServerErr(err)
	}

	return nil
}

func (repo *cardAttachmentRepository) DeleteById(
	ctx context.Context,
	id string,
) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return common.NewBadRequestErr(errors.New("invalid objectId"))
	}

	return repo.Delete(ctx, map[string]interface{}{"_id": oid})
}
