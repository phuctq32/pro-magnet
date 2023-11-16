package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	cardmodel "pro-magnet/modules/card/model"
)

func (repo *cardRepository) Create(
	ctx context.Context,
	data *cardmodel.CardCreation,
) (*cardmodel.Card, error) {
	insertData, err := data.ToCardInsert()
	if err != nil {
		return nil, err
	}

	result, err := repo.db.
		Collection(cardmodel.CardCollectionName).
		InsertOne(ctx, insertData)

	insertedId := result.InsertedID.(primitive.ObjectID).Hex()

	return &cardmodel.Card{
		Id:         &insertedId,
		CreatedAt:  insertData.CreatedAt,
		UpdatedAt:  insertData.UpdatedAt,
		Status:     insertData.Status,
		ColumnId:   data.ColumnId,
		BoardId:    data.BoardId,
		Title:      data.Title,
		Cover:      nil,
		MemberIds:  []string{},
		Checklists: []cardmodel.CardChecklist{},
		Comments:   []cardmodel.CardComment{},
		StartDate:  nil,
		EndDate:    nil,
		IsDone:     false,
	}, nil
}
