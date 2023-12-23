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

func (repo *cardChecklistRepository) Delete(
	ctx context.Context,
	cardId, checklistId string,
) error {
	cardOid, err := primitive.ObjectIDFromHex(cardId)
	if err != nil {
		return common.NewBadRequestErr(errors.New("invalid objectId"))
	}
	checklistOid, err := primitive.ObjectIDFromHex(checklistId)
	if err != nil {
		return common.NewBadRequestErr(errors.New("invalid objectId"))
	}

	filter := bson.M{"_id": cardOid}
	update := bson.M{"$pull": bson.M{"checklists": bson.M{"_id": checklistOid}}}

	if _, err := repo.db.Collection(cardchecklistmodel.CardCollectionName).
		UpdateOne(ctx, filter, update); err != nil {
		return common.NewServerErr(err)
	}

	return nil
}

func (repo *cardChecklistRepository) DeleteChecklistItem(
	ctx context.Context,
	cardId, checklistId, itemId string,
) error {
	cardOid, err := primitive.ObjectIDFromHex(cardId)
	if err != nil {
		return common.NewBadRequestErr(errors.New("invalid objectId"))
	}
	checklistOid, err := primitive.ObjectIDFromHex(checklistId)
	if err != nil {
		return common.NewBadRequestErr(errors.New("invalid objectId"))
	}
	itemOid, err := primitive.ObjectIDFromHex(itemId)
	if err != nil {
		return common.NewBadRequestErr(errors.New("invalid objectId"))
	}

	// Card filter
	filter := bson.M{"_id": cardOid}

	// Checklist filter
	identifier := []interface{}{bson.M{"checklistId._id": bson.M{"$eq": checklistOid}}}
	opts := options.Update().SetArrayFilters(options.ArrayFilters{Filters: identifier})

	update := bson.M{"$pull": bson.M{"checklists.$[checklistId].items": bson.M{"_id": itemOid}}}

	if _, err := repo.db.Collection(cardchecklistmodel.CardCollectionName).
		UpdateOne(ctx, filter, update, opts); err != nil {
		return common.NewServerErr(err)
	}

	return nil
}
