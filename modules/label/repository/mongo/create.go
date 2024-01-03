package labelrepo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"pro-magnet/common"
	labelmodel "pro-magnet/modules/label/model"
)

func (repo *labelRepository) Create(
	ctx context.Context,
	data *labelmodel.LabelCreation,
) (*labelmodel.Label, error) {
	boardOid, err := primitive.ObjectIDFromHex(data.BoardId)
	if err != nil {
		return nil, common.NewBadRequestErr(err)
	}

	insertData := &labelmodel.LabelInsert{
		Status:  labelmodel.Active,
		Title:   data.Title,
		Color:   data.Color,
		BoardId: boardOid,
	}
	result, err := repo.db.Collection(labelmodel.LabelCollectionName).InsertOne(ctx, insertData)
	if err != nil {
		return nil, common.NewServerErr(err)
	}
	insertedId := result.InsertedID.(primitive.ObjectID).Hex()

	return &labelmodel.Label{
		Id:      &insertedId,
		Title:   insertData.Title,
		Color:   insertData.Color,
		BoardId: data.BoardId,
	}, nil
}
