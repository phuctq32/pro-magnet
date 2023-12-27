package labeluc

import (
	"context"
	cardmodel "pro-magnet/modules/card/model"
	labelmodel "pro-magnet/modules/label/model"
)

type LabelRepository interface {
	Create(ctx context.Context, data *labelmodel.LabelCreation) (*labelmodel.Label, error)
	FindById(ctx context.Context, labelId string) (*labelmodel.Label, error)
	UpdateById(ctx context.Context, labelId string, updateData *labelmodel.LabelUpdate) error
	DeleteById(ctx context.Context, labelId string) error
	ExistsInBoard(ctx context.Context, labelId *string, boardId, title, color string) (bool, error)
	WithTransaction(ctx context.Context, fn func(context.Context) error) error
}

type CardRepository interface {
	UpdateLabel(ctx context.Context, cardId string, labelId string) error
	FindById(ctx context.Context, id string) (*cardmodel.Card, error)
	FindCardIdsByLabelId(ctx context.Context, labelId string) ([]string, error)
	RemoveLabel(ctx context.Context, cardId, labelId string) error
	RemoveLabelFromCardsByIds(ctx context.Context, cardIds []string, labelId string) error
}

type BoardMemberRepository interface {
	IsBoardMember(ctx context.Context, boardId, userId string) (bool, error)
}

type labelUseCase struct {
	labelRepo LabelRepository
	cardRepo  CardRepository
	bmRepo    BoardMemberRepository
}

func NewLabelUseCase(
	labelRepo LabelRepository,
	cardRepo CardRepository,
	bmRepo BoardMemberRepository,
) *labelUseCase {
	return &labelUseCase{
		labelRepo: labelRepo,
		cardRepo:  cardRepo,
		bmRepo:    bmRepo,
	}
}
