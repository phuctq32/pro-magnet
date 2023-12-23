package bmuc

import (
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"pro-magnet/common"
	bmmodel "pro-magnet/modules/boardmember/model"
)

func (uc *boardMemberUseCase) RemoveMember(
	ctx context.Context,
	requesterId string,
	data *bmmodel.BoardMember,
) error {
	board, err := uc.boardRepo.FindById(ctx, data.BoardId)
	if err != nil {
		return err
	}
	if requesterId != board.AdminId {
		return common.NewBadRequestErr(errors.New("user is not board admin"))
	}
	if data.UserId == board.AdminId {
		return common.NewBadRequestErr(errors.New("can not remove board admin"))
	}

	return uc.bmRepo.Delete(ctx, data)
}
