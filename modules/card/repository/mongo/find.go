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
	"time"
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

func (repo *cardRepository) FindByColumnId(
	ctx context.Context,
	status cardmodel.CardStatus,
	columnId string,
	cardIdsOrder []string,
) ([]cardmodel.Card, error) {
	columnOid, err := primitive.ObjectIDFromHex(columnId)
	if err != nil {
		return nil, common.NewBadRequestErr(errors.New("invalid objectId"))
	}
	cardOids := make([]primitive.ObjectID, 0)
	for _, id := range cardIdsOrder {
		oid, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return nil, common.NewBadRequestErr(errors.New("invalid objectId"))
		}
		cardOids = append(cardOids, oid)
	}

	aggPipeline := bson.A{
		bson.M{"$match": bson.M{
			"$and": bson.A{
				bson.M{"columnId": columnOid},
				bson.M{"status": status},
			},
		}},
		bson.M{"$addFields": bson.M{
			"memberCount":   bson.M{"$size": "$memberIds"},
			"commentCount":  bson.M{"$size": "$comments"},
			"orderInColumn": bson.M{"$indexOfArray": bson.A{cardOids, "$_id"}},
		}},
		bson.M{"$sort": bson.M{"orderInColumn": 1}},
		bson.M{"$project": bson.M{
			"_id":          1,
			"columnId":     1,
			"title":        1,
			"cover":        1,
			"labelIds":     1,
			"startDate":    1,
			"endDate":      1,
			"isDone":       1,
			"memberCount":  1,
			"commentCount": 1,
		}},
	}

	cursor, err := repo.db.
		Collection(cardmodel.CardCollectionName).
		Aggregate(ctx, aggPipeline)
	if err != nil {
		return nil, common.NewServerErr(err)
	}

	cards := make([]cardmodel.Card, 0)
	if err = cursor.All(ctx, &cards); err != nil {
		return nil, common.NewServerErr(err)
	}

	if cards == nil {
		return []cardmodel.Card{}, nil
	}

	return cards, err
}

func (repo *cardRepository) CountNumberOfSkillsMatchedOfEachCardDoneByUser(
	ctx context.Context, userId string, skills []string,
) ([]int, error) {
	userOid, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, common.NewBadRequestErr(errors.New("invalid objectId"))
	}
	aggPipeline := bson.A{
		bson.M{"$match": bson.M{
			"isDone":    true,
			"memberIds": userOid,
			"skills":    bson.M{"$in": skills},
		}},
		bson.M{"$project": bson.M{
			"_id": 0,
			"matchedSkillCount": bson.M{
				"$size": bson.M{
					"$filter": bson.M{
						"input": "$skills",
						"as":    "skill",
						"cond":  bson.M{"$in": bson.A{"$$skill", skills}},
					},
				},
			},
		}},
	}

	cursor, err := repo.db.Collection(cardmodel.CardCollectionName).Aggregate(ctx, aggPipeline)
	if err != nil {
		return nil, common.NewServerErr(err)
	}

	data := make([]struct {
		MatchedSkillCount int `bson:"matchedSkillCount,omitempty"`
	}, 0)
	if err = cursor.All(ctx, &data); err != nil {
		return nil, common.NewServerErr(err)
	}

	result := make([]int, len(data))
	for i := 0; i < len(data); i++ {
		result[i] = data[i].MatchedSkillCount
	}

	return result, nil
}

func (repo *cardRepository) CountCardInSamePeriodNotDoneByUser(
	ctx context.Context, userId string,
	startDate, endDate time.Time,
) (int, error) {
	userOid, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return 0, common.NewBadRequestErr(errors.New("invalid objectId"))
	}

	count, err := repo.db.
		Collection(cardmodel.CardCollectionName).
		CountDocuments(ctx, bson.M{"$and": bson.A{
			bson.M{"isDone": false},
			bson.M{"memberIds": userOid},
			bson.M{"startDate": bson.M{"$ne": nil}},
			bson.M{"endDate": bson.M{"$ne": nil}},
			bson.M{"$or": bson.A{
				bson.M{"$and": bson.A{
					bson.M{"startDate": bson.M{"$gte": startDate}},
					bson.M{"startDate": bson.M{"$lt": endDate}},
				}},
				bson.M{"$and": bson.A{
					bson.M{"endDate": bson.M{"$gt": startDate}},
					bson.M{"endDate": bson.M{"$lte": endDate}},
				}},
			}},
		}})
	if err != nil {
		return 0, common.NewServerErr(err)
	}

	return int(count), nil
}
