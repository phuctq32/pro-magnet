package recomuc

import (
	"golang.org/x/net/context"
	cardmodel "pro-magnet/modules/card/model"
	usermodel "pro-magnet/modules/user/model"
	"time"
)

type UserRepository interface {
	FindUsersByMatchingAtLeastOneCardSkills(ctx context.Context, boardId string, skills, exceptedUsedIds []string) ([]usermodel.User, error)
	WithTransaction(ctx context.Context, fn func(context.Context) error) error
}

type CardRepository interface {
	FindById(ctx context.Context, id string) (*cardmodel.Card, error)
	CountNumberOfSkillsMatchedOfEachCardDoneByUser(ctx context.Context, userId string, skills []string) ([]int, error)
	CountCardInSamePeriodNotDoneByUser(ctx context.Context, userId string, startDate, endDate time.Time) (int, error)
}

type recomUseCase struct {
	userRepo UserRepository
	cardRepo CardRepository
}

func NewRecommendationUseCase(
	userRepo UserRepository,
	cardRepo CardRepository,
) *recomUseCase {
	return &recomUseCase{
		userRepo: userRepo,
		cardRepo: cardRepo,
	}
}
