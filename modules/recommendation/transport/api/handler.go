package recomapi

import (
	"golang.org/x/net/context"
	usermodel "pro-magnet/modules/user/model"
)

type RecommendationUseCase interface {
	GetRecommendedUsersForCard(ctx context.Context, requesterId, cardId string, quantity int) ([]usermodel.User, error)
}

type recomHandler struct {
	uc RecommendationUseCase
}

func NewRecommendationHandler(uc RecommendationUseCase) *recomHandler {
	return &recomHandler{uc: uc}
}
