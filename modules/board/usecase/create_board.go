package boarduc

import (
	"context"
	boardmodel "pro-magnet/modules/board/model"
	"slices"
)

func (uc *boardUseCase) CreateBoard(
	ctx context.Context,
	data *boardmodel.BoardCreation,
) (*boardmodel.Board, error) {
	errCh := make(chan error, 2)

	go func() {
		// Board name must be not existed
		ok, err := uc.boardRepo.ExistsInWorkspace(ctx, data.WorkspaceId)
		if err != nil {
			errCh <- err
		} else if ok {
			errCh <- boardmodel.ErrExistedBoard
		} else {
			errCh <- nil
		}
	}()

	go func() {
		// Check if board admin is a workspace member
		wsMemberIds, err := uc.wsRepo.GetMemberIds(ctx, data.WorkspaceId)
		if err != nil {
			errCh <- err
			return
		}
		if !slices.Contains(wsMemberIds, data.UserId) {
			errCh <- boardmodel.ErrIsNotMemberOfWorkspace
			return
		}
		errCh <- nil
	}()

	var err error
	for i := 0; i < 2; i++ {
		e := <-errCh
		if err != nil {
			continue
		} else if e != nil {
			err = e
		}
	}
	close(errCh)
	if err != nil {
		return nil, err
	}

	return uc.boardRepo.Create(ctx, data)
}
