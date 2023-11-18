package boardmemberrepo

import (
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/net/context"
	"pro-magnet/common"
	"pro-magnet/modules/boardmember/model"
)

func (repo *boardMemberRepository) IsBoardMember(
	ctx context.Context,
	boardId, userId string,
) (bool, error) {
	boardOid, err := primitive.ObjectIDFromHex(boardId)
	if err != nil {
		return false, common.NewBadRequestErr(errors.New("invalid objectId"))
	}

	userOid, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return false, common.NewBadRequestErr(errors.New("invalid objectId"))
	}

	count, err := repo.db.
		Collection(model.BoardMemberCollectionName).
		CountDocuments(ctx, bson.M{
			"boardId": boardOid,
			"userId":  userOid,
		})
	if err != nil {
		return false, common.NewServerErr(err)
	}

	return count > 0, nil
}
