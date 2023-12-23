package boardrepo

import (
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
	"pro-magnet/common"
	boardmodel "pro-magnet/modules/board/model"
)

func (repo *boardRepository) FindOne(
	ctx context.Context,
	filter map[string]interface{},
) (*boardmodel.Board, error) {
	var board boardmodel.Board

	if err := repo.db.
		Collection(boardmodel.BoardCollectionName).
		FindOne(ctx, filter).Decode(&board); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, common.NewNotFoundErr("board", boardmodel.ErrBoardNotFound)
		}

		return nil, common.NewServerErr(err)
	}

	return &board, nil
}

func (repo *boardRepository) FindById(
	ctx context.Context,
	id string,
) (*boardmodel.Board, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	return repo.FindOne(ctx, map[string]interface{}{"_id": oid})
}
