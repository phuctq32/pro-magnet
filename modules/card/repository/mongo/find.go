package mongo

import (
	"context"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (repo *cardRepository) FindCardIdsByLabelId(
	ctx context.Context,
	labelId string,
) ([]string, error) {
	labelOid, err := primitive.ObjectIDFromHex(labelId)
	if err != nil {
		return nil, common.NewBadRequestErr(err)
	}

	opt := options.Find().SetProjection(bson.M{
		"_id": 1,
	})
	cursor, err := repo.db.
		Collection(cardmodel.CardCollectionName).
		Find(ctx, bson.M{"labelIds": labelOid}, opt)
	if err != nil {
		return nil, common.NewServerErr(err)
	}

	var cards []cardmodel.Card
	if err = cursor.All(ctx, &cards); err != nil {
		return nil, common.NewServerErr(err)
	}

	if cards == nil {
		return []string{}, nil
	}

	res := make([]string, 0)
	for i := 0; i < len(cards); i++ {
		res = append(res, *cards[i].Id)
	}

	return res, nil
}
