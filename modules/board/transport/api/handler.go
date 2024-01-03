package boardapi

import (
	"context"
	boardmodel "pro-magnet/modules/board/model"
)

type BoardUseCase interface {
	CreateBoard(ctx context.Context, data *boardmodel.BoardCreation) (*boardmodel.Board, error)
	UpdateBoard(ctx context.Context, requesterId, boardId string, updateData *boardmodel.BoardUpdate) error
	GetBoardById(ctx context.Context, requesterId, boardId string) (board *boardmodel.Board, err error)
}

type boardHandler struct {
	uc BoardUseCase
}

func NewBoardHandler(uc BoardUseCase) *boardHandler {
	return &boardHandler{uc: uc}
}
