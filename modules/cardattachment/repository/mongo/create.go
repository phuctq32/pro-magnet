package carepo

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/net/context"
	"pro-magnet/common"
	camodel "pro-magnet/modules/cardattachment/model"
)

func (repo *cardAttachmentRepository) Create(
	ctx context.Context,
	data *camodel.CardAttachment,
) (*camodel.CardAttachment, error) {
	dataInsert, err := data.ToDataInsert()
	if err != nil {
		return nil, err
	}

	result, err := repo.db.
		Collection(camodel.CardAttachmentCollectionName).
		InsertOne(ctx, dataInsert)
	if err != nil {
		return nil, common.NewServerErr(err)
	}
	insertedId := result.InsertedID.(primitive.ObjectID).Hex()
	data.Id = &insertedId
	data.CreatedAt = dataInsert.CreatedAt

	return data, nil
}
