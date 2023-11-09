package mongo

import (
	"context"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"pro-magnet/common"
	cardmodel "pro-magnet/modules/card/model"
)

func (repo *cardRepository) FindOne(
	ctx context.Context,
	filter map[string]interface{},
) (*cardmodel.Card, error) {
	var card cardmodel.Card

	if err := repo.db.
		Collection(cardmodel.CardCollectionName).
		FindOne(ctx, filter).Decode(&card); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, common.NewNotFoundErr("card", cardmodel.ErrCardNotFound)
		}

		return nil, common.NewServerErr(err)
	}

	return &card, nil
}

func (repo *cardRepository) FindById(
	ctx context.Context,
	id string,
) (*cardmodel.Card, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	return repo.FindOne(ctx, map[string]interface{}{"_id": oid})
}
