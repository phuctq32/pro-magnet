package cardchecklistrepo

import (
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
