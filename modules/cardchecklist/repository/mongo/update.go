package cardchecklistrepo

import (
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
	"pro-magnet/common"
	cardchecklistmodel "pro-magnet/modules/cardchecklist/model"
)

func (repo *cardChecklistRepository) Update(
	ctx context.Context,
	cardId, checklistId string,
	data *cardchecklistmodel.CardChecklistUpdate,
) error {
	cardOid, err := primitive.ObjectIDFromHex(cardId)
	if err != nil {
		return common.NewBadRequestErr(errors.New("invalid objectId"))
	}
	checklistOid, err := primitive.ObjectIDFromHex(checklistId)
	if err != nil {
		return common.NewBadRequestErr(errors.New("invalid objectId"))
	}

	// Card filter
	filter := bson.M{"_id": cardOid}

	// Checklist filter
	identifier := []interface{}{bson.M{"checklist._id": bson.M{"$eq": checklistOid}}}
	opts := options.Update().
		SetArrayFilters(options.ArrayFilters{Filters: identifier})

	update := bson.M{"$set": bson.M{"checklists.$[checklist].name": data.Name}}

	if _, err = repo.db.Collection(cardchecklistmodel.CardCollectionName).
		UpdateOne(ctx, filter, update, opts); err != nil {
		return common.NewServerErr(err)
	}

	return nil
}
