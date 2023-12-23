package boardrepo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"pro-magnet/common"
	boardmodel "pro-magnet/modules/board/model"
	"time"
)

func (repo *boardRepository) Create(
	ctx context.Context,
	data *boardmodel.BoardCreation,
) (*boardmodel.Board, error) {
	wsOid, _ := primitive.ObjectIDFromHex(data.WorkspaceId)
	adminOid, _ := primitive.ObjectIDFromHex(data.UserId)
	now := time.Now()

	boardInsert := &boardmodel.BoardInsert{
		CreatedAt:      now,
		UpdatedAt:      now,
		Name:           data.Name,
		WorkspaceId:    wsOid,
		AdminId:        adminOid,
		ColumnOrderIds: []primitive.ObjectID{},
	}

	result, err := repo.db.
		Collection(boardmodel.BoardCollectionName).
		InsertOne(ctx, boardInsert)
	if err != nil {
		return nil, common.NewServerErr(err)
	}

	insertedId := result.InsertedID.(primitive.ObjectID).Hex()

	return &boardmodel.Board{
		Id:             &insertedId,
		CreatedAt:      boardInsert.CreatedAt,
		UpdatedAt:      boardInsert.UpdatedAt,
		Name:           boardInsert.Name,
		WorkspaceId:    data.WorkspaceId,
		AdminId:        data.UserId,
		ColumnOrderIds: []string{},
	}, nil
}
