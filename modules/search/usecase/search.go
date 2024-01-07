package searchuc

import (
	"golang.org/x/net/context"
	searchmodel "pro-magnet/modules/search/model"
)

func (uc *searchUseCase) Search(
	ctx context.Context,
	requesterId, searchTerm string,
) (data *searchmodel.SearchData, err error) {
	data = new(searchmodel.SearchData)
	err = uc.asyncg.Process(
		ctx,
		uc.searchWorkspaces(data, requesterId, searchTerm),
		uc.searchBoards(data, requesterId, searchTerm),
		uc.searchCards(data, requesterId, searchTerm),
	)
	if err != nil {
		return nil, err
	}

	return
}

func (uc *searchUseCase) searchWorkspaces(
	data *searchmodel.SearchData,
	memberId, searchTerm string,
) func(context.Context) error {
	return func(ctx context.Context) error {
		workspaces, err := uc.wsRepo.Search(ctx, memberId, searchTerm)
		if err != nil {
			return err
		}

		data.Workspaces = workspaces
		return nil
	}
}

func (uc *searchUseCase) searchBoards(
	data *searchmodel.SearchData,
	memberId, searchTerm string,
) func(context.Context) error {
	return func(ctx context.Context) error {
		boards, err := uc.boardRepo.Search(ctx, memberId, searchTerm)
		if err != nil {
			return err
		}

		data.Boards = boards
		return nil
	}
}

func (uc *searchUseCase) searchCards(
	data *searchmodel.SearchData,
	memberId, searchTerm string,
) func(context.Context) error {
	return func(ctx context.Context) error {
		cards, err := uc.careRepo.Search(ctx, memberId, searchTerm)
		if err != nil {
			return err
		}

		data.Cards = cards
		return nil
	}
}
