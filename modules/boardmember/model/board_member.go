package bmmodel

import (
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"pro-magnet/common"
)

type BoardMemberInsert struct {
	BoardId primitive.ObjectID `bson:"boardId"`
	UserId  primitive.ObjectID `bson:"userId"`
}

type BoardMember struct {
	BoardId string `json:"boardId" bson:"boardId" validate:"required,mongodb"`
	UserId  string `json:"userId" bson:"userId" validate:"required,mongodb"`
}

func (bm *BoardMember) ToBoardMemberInsert() (*BoardMemberInsert, error) {
	boardOid, err := primitive.ObjectIDFromHex(bm.BoardId)
	if err != nil {
		return nil, common.NewBadRequestErr(errors.New("invalid objectId"))
	}

	userOid, err := primitive.ObjectIDFromHex(bm.UserId)
	if err != nil {
		return nil, common.NewBadRequestErr(errors.New("invalid objectId"))
	}

	return &BoardMemberInsert{
		BoardId: boardOid,
		UserId:  userOid,
	}, nil
}
