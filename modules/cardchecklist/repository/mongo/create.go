package cardchecklistrepo

import (
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
	"pro-magnet/common"
	cardmodel "pro-magnet/modules/card/model"
	cardchecklistmodel "pro-magnet/modules/cardchecklist/model"
)

func (repo *cardChecklistRepository) Create(
	ctx context.Context,
	cardId string,
	data *cardchecklistmodel.CardChecklist,
) error {
	cardOid, err := primitive.ObjectIDFromHex(cardId)
	if err != nil {
		return common.NewBadRequestErr(errors.New("invalid objectId"))
	}

	insertData := &cardchecklistmodel.CardChecklistInsert{
		Id:    primitive.NewObjectID(),
		Name:  data.Name,
		Items: data.Items,
	}

	filter := bson.M{"_id": cardOid}
	update := bson.M{"$push": bson.M{"checklists": insertData}}
	result, err := repo.db.
		Collection(cardchecklistmodel.CardCollectionName).
		UpdateOne(ctx, filter, update)
	if err != nil {
		return common.NewServerErr(err)
	}
	if result.MatchedCount == 0 {
		return common.NewBadRequestErr(cardmodel.ErrCardNotFound)
	}

	return nil
}

func (repo *cardChecklistRepository) CreateChecklistItem(
	ctx context.Context,
	cardId, checklistId string,
	data *cardchecklistmodel.ChecklistItem,
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
	opts := options.Update().SetArrayFilters(options.ArrayFilters{Filters: identifier})

	// Add items to checklist
	insertData := &cardchecklistmodel.ChecklistItemInsert{
		Id:     primitive.NewObjectID(),
		Title:  data.Title,
		IsDone: false,
	}
	update := bson.M{"$push": bson.M{"checklists.$[checklist].items": insertData}}

	if _, err := repo.db.Collection(cardchecklistmodel.CardCollectionName).
		UpdateOne(ctx, filter, update, opts); err != nil {
		return common.NewServerErr(err)
	}

	return nil
}
