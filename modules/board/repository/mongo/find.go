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

func (repo *boardRepository) Find(
	ctx context.Context,
	status boardmodel.BoardStatus,
	filter map[string]interface{},
) ([]boardmodel.Board, error) {
	filter["status"] = status
	cursor, err := repo.db.
		Collection(boardmodel.BoardCollectionName).
		Find(ctx, filter)
	if err != nil {
		return nil, common.NewServerErr(err)
	}

	result := make([]boardmodel.Board, 0)
	if err = cursor.All(ctx, &result); err != nil {
		return nil, common.NewServerErr(err)
	}

	return result, nil
}

func (repo *boardRepository) FindByWorkspaceId(
	ctx context.Context,
	status boardmodel.BoardStatus,
	workspaceId string,
) ([]boardmodel.Board, error) {
	wsOid, err := primitive.ObjectIDFromHex(workspaceId)
	if err != nil {
		return nil, common.NewBadRequestErr(errors.New("invalid objectId"))
	}

	return repo.Find(ctx, status, map[string]interface{}{"workspaceId": wsOid})
}
