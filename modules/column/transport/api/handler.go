package columnapi

import (
	"golang.org/x/net/context"
	columnmodel "pro-magnet/modules/column/model"
)

type ColumnUseCase interface {
	CreateColumn(ctx context.Context, userId string, data *columnmodel.ColumnCreate) (*columnmodel.Column, error)
	UpdateColumn(ctx context.Context, userId, columnId string, data *columnmodel.ColumnUpdate) (*columnmodel.Column, error)
	RemoveColumn(ctx context.Context, userId string, columnId string) error
}

type columnHandler struct {
	uc ColumnUseCase
}

func NewColumnHandler(uc ColumnUseCase) *columnHandler {
	return &columnHandler{uc: uc}
}
