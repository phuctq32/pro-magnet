package mongo

import (
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/net/context"
	"pro-magnet/common"
	cardmodel "pro-magnet/modules/card/model"
)

func (repo *cardRepository) Delete(
	ctx context.Context,
	filter map[string]interface{},
) error {
	_, err := repo.db.
		Collection(cardmodel.CardCollectionName).
		UpdateMany(ctx, filter, bson.M{
			"$set": bson.M{
				"status": cardmodel.Deleted,
			},
			"$currentDate": bson.M{
				"updatedAt": bson.M{"$type": "date"},
			},
		})
	if err != nil {
		return common.NewServerErr(err)
	}

	return nil
}

func (repo *cardRepository) DeleteById(
	ctx context.Context,
	id string,
) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return common.NewBadRequestErr(errors.New("invalid object id"))
	}

	return repo.Delete(ctx, map[string]interface{}{"_id": oid})
}

func (repo *cardRepository) DeleteByIds(
	ctx context.Context,
	ids []string,
) error {
	oids := make([]primitive.ObjectID, len(ids))
	for i, id := range ids {
		oid, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return common.NewBadRequestErr(errors.New("invalid object id"))
		}

		oids[i] = oid
	}

	return repo.Delete(ctx, bson.M{"_id": bson.M{"$in": oids}})
}
