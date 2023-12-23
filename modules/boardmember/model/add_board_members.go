package bmmodel

import (
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"pro-magnet/common"
)

type AddBoardMembers struct {
	UserIds []string `json:"userIds" validate:"required,min=1,dive,mongodb"`
	BoardId string   `json:"boardId" validate:"required,mongodb"`
}

func (bm *AddBoardMembers) ToBoardMembersInsert() ([]interface{}, error) {
	boardOid, err := primitive.ObjectIDFromHex(bm.BoardId)
	if err != nil {
		return nil, common.NewBadRequestErr(errors.New("invalid objectId"))
	}

	res := make([]interface{}, 0)
	for i := 0; i < len(bm.UserIds); i++ {
		userOid, err := primitive.ObjectIDFromHex(bm.UserIds[i])
		if err != nil {
			return nil, common.NewBadRequestErr(errors.New("invalid objectId"))
		}
		res = append(res, bson.M{"boardId": boardOid, "userId": userOid})
	}

	return res, nil
}
