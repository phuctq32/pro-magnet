package cardcommentuc

import (
	"golang.org/x/net/context"
	cardmodel "pro-magnet/modules/card/model"
	cardcommentmodel "pro-magnet/modules/cardcomment/model"
)

type BoardMemberRepository interface {
	IsBoardMember(ctx context.Context, boardId, userId string) (bool, error)
}

type CardRepository interface {
	FindById(ctx context.Context, id string) (*cardmodel.Card, error)
}

type CardCommentRepository interface {
	Create(ctx context.Context, cardId string, data *cardcommentmodel.CardCommentCreate) error
	Update(ctx context.Context, cardId, commentId string, updateData *cardcommentmodel.CardCommentUpdate) error
	Delete(ctx context.Context, cardId, commentId string) error
}

type cardCommentUseCase struct {
	cmRepo   CardCommentRepository
	cardRepo CardRepository
	bmRepo   BoardMemberRepository
}

func NewCardCommentUseCase(
	cmRepo CardCommentRepository,
	cardRepo CardRepository,
	bmRepo BoardMemberRepository,
) *cardCommentUseCase {
	return &cardCommentUseCase{
		cmRepo:   cmRepo,
		cardRepo: cardRepo,
		bmRepo:   bmRepo,
	}
}
