package columnrepo

import (
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/net/context"
	"pro-magnet/common"
	columnmodel "pro-magnet/modules/column/model"
)

func (repo *columnRepository) FindByBoardId(
	ctx context.Context,
	status columnmodel.ColumnStatus,
	boardId string,
	columnIdsOrder []string,
) ([]columnmodel.Column, error) {
	boardOid, err := primitive.ObjectIDFromHex(boardId)
	if err != nil {
		return nil, common.NewBadRequestErr(errors.New("invalid objectId"))
	}
	colOids := make([]primitive.ObjectID, 0)
	for _, id := range columnIdsOrder {
		oid, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return nil, common.NewBadRequestErr(errors.New("invalid objectId"))
		}
		colOids = append(colOids, oid)
	}

	aggPipeline := bson.A{
		bson.M{"$match": bson.M{
			"$and": bson.A{
				bson.M{"boardId": boardOid},
				bson.M{"status": status},
			},
		}},
		bson.M{"$addFields": bson.M{
			"orderInBoard": bson.M{"$indexOfArray": bson.A{colOids, "$_id"}},
		}},
		bson.M{"$sort": bson.M{"orderInBoard": 1}},
		bson.M{"$project": bson.M{"orderInBoard": 0}},
	}
	cursor, err := repo.db.
		Collection(columnmodel.ColumnCollectionName).
		Aggregate(ctx, aggPipeline)
	if err != nil {
		return nil, common.NewServerErr(err)
	}

	cols := make([]columnmodel.Column, 0)
	if err = cursor.All(ctx, &cols); err != nil {
		return nil, common.NewServerErr(err)
	}

	return cols, err
}
