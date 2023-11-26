package columnuc

import (
	"golang.org/x/net/context"
	columnmodel "pro-magnet/modules/column/model"
)

type BoardMemberRepository interface {
	IsBoardMember(ctx context.Context, boardId, userId string) (bool, error)
}

type CardRepository interface {
	DeleteByIds(ctx context.Context, ids []string) error
}

type ColumnRepository interface {
	Create(ctx context.Context, data *columnmodel.ColumnCreate) (*columnmodel.Column, error)
	DeleteById(ctx context.Context, id string) error
	FindById(ctx context.Context, id string) (*columnmodel.Column, error)
	UpdateById(ctx context.Context, id string, updateData *columnmodel.ColumnUpdate) (*columnmodel.Column, error)
	WithTransaction(ctx context.Context, fn func(context.Context) error) error
}

type columnUseCase struct {
	colRepo  ColumnRepository
	bmRepo   BoardMemberRepository
	cardRepo CardRepository
}

func NewColumnUseCase(
	colRepo ColumnRepository,
	bmRepo BoardMemberRepository,
	cardRepo CardRepository,
) *columnUseCase {
	return &columnUseCase{
		colRepo:  colRepo,
		bmRepo:   bmRepo,
		cardRepo: cardRepo,
	}
}
