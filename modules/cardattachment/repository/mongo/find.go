package carepo

import (
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
	"pro-magnet/common"
	camodel "pro-magnet/modules/cardattachment/model"
)

func (repo *cardAttachmentRepository) FindOne(
	ctx context.Context,
	filter map[string]interface{},
) (*camodel.CardAttachment, error) {
	var res camodel.CardAttachment

	if err := repo.db.
		Collection(camodel.CardAttachmentCollectionName).
		FindOne(ctx, filter).Decode(&res); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, common.NewNotFoundErr("card attachment", err)
		}

		return nil, common.NewServerErr(err)
	}

	return &res, nil
}

func (repo *cardAttachmentRepository) FindById(
	ctx context.Context,
	id string,
) (*camodel.CardAttachment, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, common.NewBadRequestErr(errors.New("invalid objectId"))
	}

	return repo.FindOne(ctx, map[string]interface{}{"_id": oid})
}

func (repo *cardAttachmentRepository) CountByCardId(
	ctx context.Context,
	caStatus camodel.CardAttachmentStatus,
	cardId string,
) (int, error) {
	cardOid, err := primitive.ObjectIDFromHex(cardId)
	if err != nil {
		return 0, common.NewBadRequestErr(errors.New("invalid objectId"))
	}

	count, err := repo.db.
		Collection(camodel.CardAttachmentCollectionName).
		CountDocuments(ctx, bson.M{
			"cardId": cardOid,
			"status": caStatus,
		})
	if err != nil {
		return 0, common.NewServerErr(err)
	}

	return int(count), nil
}
