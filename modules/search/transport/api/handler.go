package searchapi

import (
	"golang.org/x/net/context"
	searchmodel "pro-magnet/modules/search/model"
)

type SearchUseCase interface {
	Search(ctx context.Context, requesterId, searchStr string) (*searchmodel.SearchData, error)
}

type searchHandler struct {
	uc SearchUseCase
}

func NewSearchHandler(uc SearchUseCase) *searchHandler {
	return &searchHandler{uc: uc}
}
