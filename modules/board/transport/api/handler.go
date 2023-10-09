package boardapi

import (
	"context"
	boardmodel "pro-magnet/modules/board/model"
)

type BoardUseCase interface {
	CreateBoard(ctx context.Context, data *boardmodel.BoardCreation) (*boardmodel.Board, error)
}

type boardHandler struct {
	uc BoardUseCase
}

func NewBoardHandler(uc BoardUseCase) *boardHandler {
	return &boardHandler{uc: uc}
}
