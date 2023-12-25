package bmuc

import (
	"golang.org/x/net/context"
	"pro-magnet/components/asyncgroup"
	boardmodel "pro-magnet/modules/board/model"
	bmmodel "pro-magnet/modules/boardmember/model"
	usermodel "pro-magnet/modules/user/model"
)

type UserRepository interface {
	FindSimpleUsersByIds(ctx context.Context, userIds []string) ([]usermodel.User, error)
}

type BoardRepository interface {
	FindById(ctx context.Context, id string) (*boardmodel.Board, error)
}

type BoardMemberRepository interface {
	IsBoardMember(ctx context.Context, boardId, userId string) (bool, error)
	CreateMany(ctx context.Context, data *bmmodel.AddBoardMembers) error
	Delete(ctx context.Context, data *bmmodel.BoardMember) error
	FindMemberIdsByBoardId(ctx context.Context, boardId string) ([]string, error)
}

type boardMemberUseCase struct {
	bmRepo    BoardMemberRepository
	boardRepo BoardRepository
	userRepo  UserRepository
	asyncg    asyncgroup.AsyncGroup
}

func NewBoardMemberUseCase(
	bmRepo BoardMemberRepository,
	boardRepo BoardRepository,
	userRepo UserRepository,
	asyncg asyncgroup.AsyncGroup,
) *boardMemberUseCase {
	return &boardMemberUseCase{
		bmRepo:    bmRepo,
		boardRepo: boardRepo,
		userRepo:  userRepo,
		asyncg:    asyncg,
	}
}
