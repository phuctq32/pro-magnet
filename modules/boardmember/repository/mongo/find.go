package boardmemberrepo

import (
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/net/context"
	"pro-magnet/common"
	bmmodel "pro-magnet/modules/boardmember/model"
)

func (repo *boardMemberRepository) FindMemberIdByBoardId(
	ctx context.Context,
	boardId string,
) ([]string, error) {
	boardOid, err := primitive.ObjectIDFromHex(boardId)
	if err != nil {
		return nil, common.NewBadRequestErr(errors.New("invalid objectId"))
	}

	cursor, err := repo.db.
		Collection(bmmodel.BoardMemberCollectionName).
		Find(ctx, bson.M{"boardId": boardOid})
	if err != nil {
		return nil, common.NewServerErr(err)
	}

	var bms []bmmodel.BoardMember
	if err = cursor.All(ctx, &bms); err != nil {
		return nil, common.NewServerErr(err)
	}

	if bms == nil {
		return []string{}, nil
	}

	result := make([]string, 0)
	for i := 0; i < len(bms); i++ {
		result = append(result, bms[i].UserId)
	}

	return result, nil
}
