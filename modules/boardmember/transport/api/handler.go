package bmapi

import (
	"golang.org/x/net/context"
	bmmodel "pro-magnet/modules/boardmember/model"
	usermodel "pro-magnet/modules/user/model"
)

type BoardMemberUseCase interface {
	AddMember(ctx context.Context, requesterId string, data *bmmodel.AddBoardMembers) error
	RemoveMember(ctx context.Context, requesterId string, data *bmmodel.BoardMember) error
	GetBoardMembers(ctx context.Context, requesterId string, boardId string) ([]usermodel.User, error)
}

type boardMemberHandler struct {
	uc BoardMemberUseCase
}

func NewBoardMemberHandler(uc BoardMemberUseCase) *boardMemberHandler {
	return &boardMemberHandler{uc: uc}
}
