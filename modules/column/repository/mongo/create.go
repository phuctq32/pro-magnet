package columnrepo

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/net/context"
	columnmodel "pro-magnet/modules/column/model"
)

func (repo *columnRepository) Create(
	ctx context.Context,
	data *columnmodel.ColumnCreate,
) (*columnmodel.Column, error) {
	insertData, err := data.ToColumnInsert()
	if err != nil {
		return nil, err
	}

	result, err := repo.db.Collection(columnmodel.ColumnCollectionName).InsertOne(ctx, insertData)
	if err != nil {
		return nil, err
	}
	insertedId := result.InsertedID.(primitive.ObjectID).Hex()

	return &columnmodel.Column{
		Id:             &insertedId,
		Status:         data.Status,
		Title:          data.Title,
		BoardId:        data.BoardId,
		OrderedCardIds: make([]string, 0),
	}, nil
}
