package bmuc

import (
	"golang.org/x/net/context"
	bmmodel "pro-magnet/modules/boardmember/model"
)

type BoardMemberRepository interface {
	IsBoardMember(ctx context.Context, boardId, userId string) (bool, error)
	Create(ctx context.Context, data *bmmodel.BoardMember) error
}

type boardMemberUseCase struct {
	bmRepo BoardMemberRepository
}

func NewBoardMemberUseCase(
	bmRepo BoardMemberRepository,
) *boardMemberUseCase {
	return &boardMemberUseCase{
		bmRepo: bmRepo,
	}
}
