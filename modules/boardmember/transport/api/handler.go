package bmapi

import (
	"golang.org/x/net/context"
	bmmodel "pro-magnet/modules/boardmember/model"
)

type BoardMemberUseCase interface {
	AddMember(ctx context.Context, requesterId string, data *bmmodel.BoardMember) error
}

type boardMemberHandler struct {
	uc BoardMemberUseCase
}

func NewBoardMemberHandler(uc BoardMemberUseCase) *boardMemberHandler {
	return &boardMemberHandler{uc: uc}
}
