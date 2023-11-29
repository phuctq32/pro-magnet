package carepo

import (
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/net/context"
	"pro-magnet/common"
	camodel "pro-magnet/modules/cardattachment/model"
)

func (repo *cardAttachmentRepository) List(
	ctx context.Context,
	status camodel.CardAttachmentStatus,
	filter map[string]interface{},
) ([]camodel.CardAttachment, error) {
	filter["status"] = status
	cursor, err := repo.db.
		Collection(camodel.CardAttachmentCollectionName).
		Find(ctx, filter)
	if err != nil {
		return nil, common.NewServerErr(err)
	}

	result := make([]camodel.CardAttachment, 0)
	if err = cursor.All(ctx, &result); err != nil {
		return nil, common.NewServerErr(err)
	}

	return result, nil
}

func (repo *cardAttachmentRepository) ListByCardId(
	ctx context.Context,
	status camodel.CardAttachmentStatus,
	cardId string,
) ([]camodel.CardAttachment, error) {
	cardOid, err := primitive.ObjectIDFromHex(cardId)
	if err != nil {
		return nil, common.NewBadRequestErr(errors.New("invalid object id"))
	}

	return repo.List(ctx, status, map[string]interface{}{
		"cardId": cardOid,
	})
}
