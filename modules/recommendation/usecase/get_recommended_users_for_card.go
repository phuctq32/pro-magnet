package recomuc

import (
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"golang.org/x/net/context"
	"pro-magnet/common"
	cardmodel "pro-magnet/modules/card/model"
	usermodel "pro-magnet/modules/user/model"
	"sort"
)

func (uc *recomUseCase) GetRecommendedUsersForCard(
	ctx context.Context,
	requesterId, cardId string,
	quantity int,
) (users []usermodel.User, err error) {
	err = uc.userRepo.WithTransaction(ctx, func(txCtx context.Context) error {
		card, e := uc.cardRepo.FindById(txCtx, cardId)
		if e != nil {
			return e
		}
		if card.Status == cardmodel.Deleted {
			return common.NewBadRequestErr(cardmodel.ErrCardDeleted)
		}
		if card.IsDone {
			return common.NewBadRequestErr(errors.New("cannot recommend users for card is done"))
		}
		if card.Skills == nil || len(card.Skills) == 0 {
			users = make([]usermodel.User, 0)
			return nil
		}

		requiredSkillMap := make(map[string]bool)
		for _, skill := range card.Skills {
			requiredSkillMap[skill] = true
		}

		users, e = uc.userRepo.FindUsersByMatchingAtLeastOneCardSkills(txCtx, *card.BoardId, card.Skills, card.MemberIds)
		if err != nil {
			return err
		}

		for i := 0; i < len(users); i++ {
			// For each skill of user matched required skills, increase score by 1
			for _, skill := range users[i].Skills {
				if requiredSkillMap[skill] {
					users[i].SkillScore += 1
				}
			}

			// Count the number of skills of each card which has at least 1 skill matched required skill
			// And done by user.
			// Then for each card, the first matched skill score is 0.005, the follows score is 0.01
			// E.g: required skills is ["English", "Java"]
			// Cards has data like:
			// [
			// {...another field, skills: ["Java"], memberdIds: [..., userId]}
			// {...another field, skills: ["English", "Java"], memberdIds: [..., userId]}
			//]
			// The result is data = [ 1, 2 ]
			// Score will be calculated: (1*0.005) + ((2-1)*0.01 + 1*0.005) = 0.02
			data, e := uc.cardRepo.CountNumberOfSkillsMatchedOfEachCardDoneByUser(txCtx, *users[i].Id, card.Skills)
			if e != nil {
				return e
			}
			for _, val := range data {
				users[i].SkillScore += float32(val-1)*0.01 + 0.005
			}

			// Count cards is not done and startDate, endDate in same period
			// For each card, decrease score by 0.1
			if card.StartDate != nil && card.EndDate != nil {
				joinedCardCount, e := uc.cardRepo.CountCardInSamePeriodNotDoneByUser(txCtx, *users[i].Id, *card.StartDate, *card.EndDate)
				if e != nil {
					return e
				}
				users[i].SkillScore -= (0.1 * float32(joinedCardCount))
			}

			log.Debug().Str("email", *users[i].Email).Float32("score", users[i].SkillScore).Msg("")
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	sort.Slice(users, func(i, j int) bool {
		return users[i].SkillScore > users[j].SkillScore
	})
	if len(users) > quantity {
		return users[:quantity], nil
	}

	return users, nil
}
