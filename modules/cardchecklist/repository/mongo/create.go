package cardchecklistrepo

import (
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
