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

func (repo *cardRepository) UpdateById(
	ctx context.Context,
	id string,
	updateData *cardmodel.CardUpdate,
) (*cardmodel.Card, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	return repo.UpdateOne(ctx, map[string]interface{}{"_id": oid}, updateData)
}

func (repo *cardRepository) UpdateOne(
	ctx context.Context,
	filter map[string]interface{},
	updateData *cardmodel.CardUpdate,
) (*cardmodel.Card, error) {
	return updateCard(ctx, repo.db, filter, updateData)
}

func updateCard[T map[string]interface{} | *cardmodel.CardUpdate](
	ctx context.Context,
	db *mongo.Database,
	filter map[string]interface{},
	updateData T,
) (*cardmodel.Card, error) {
	var updatedCard cardmodel.Card

	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	if err := db.
		Collection(cardmodel.CardCollectionName).
		FindOneAndUpdate(ctx, filter, bson.M{
			"$set": updateData,
			"$currentDate": bson.M{
				"updatedAt": bson.M{"$type": "date"},
			},
		}, opts).Decode(&updatedCard); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, common.NewNotFoundErr("card", cardmodel.ErrCardNotFound)
		}
		return nil, common.NewServerErr(err)
	}

	return &updatedCard, nil
}
