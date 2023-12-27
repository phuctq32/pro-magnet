package boardrepo

import (
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
	"pro-magnet/common"
	boardmodel "pro-magnet/modules/board/model"
)

func (repo *boardRepository) Update(
	ctx context.Context,
	filter map[string]interface{},
	updateData *boardmodel.BoardUpdate,
) error {
	result, err := repo.db.
		Collection(boardmodel.BoardCollectionName).
		UpdateOne(ctx, filter, bson.M{"$set": updateData})
	if err != nil {
		return common.NewServerErr(err)
	}
	if result.MatchedCount == 0 {
		return common.NewNotFoundErr("board", mongo.ErrNoDocuments)
	}

	return nil
}

func (repo *boardRepository) UpdateById(
	ctx context.Context,
	boardId string,
	updateData *boardmodel.BoardUpdate,
) error {
	boardOid, err := primitive.ObjectIDFromHex(boardId)
	if err != nil {
		return common.NewBadRequestErr(errors.New("invalid objectId"))
	}

	return repo.Update(ctx, map[string]interface{}{"_id": boardOid}, updateData)
}

func (repo *boardRepository) AddColumnId(
	ctx context.Context,
	boardId, columnId string,
) error {
	boardOid, err := primitive.ObjectIDFromHex(boardId)
	if err != nil {
		return common.NewBadRequestErr(errors.New("invalid objectId"))
	}
	columnOid, err := primitive.ObjectIDFromHex(columnId)
	if err != nil {
		return common.NewBadRequestErr(errors.New("invalid objectId"))
	}

	result, err := repo.db.
		Collection(boardmodel.BoardCollectionName).
		UpdateOne(
			ctx,
			bson.M{"_id": boardOid},
			bson.M{"$push": bson.M{"orderedColumnIds": columnOid}},
		)
	if err != nil {
		return common.NewServerErr(err)
	}
	if result.MatchedCount == 0 {
		return common.NewNotFoundErr("board", mongo.ErrNoDocuments)
	}

	return nil
}
