package columnmodel

import (
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"pro-magnet/common"
)

type ColumnInsert struct {
	Status         ColumnStatus         `bson:"status"`
	Title          string               `bson:"title"`
	BoardId        primitive.ObjectID   `bson:"boardId"`
	OrderedCardIds []primitive.ObjectID `bson:"orderedCardIds"`
}

type ColumnCreate struct {
	Title   string       `json:"title" validate:"required"`
	BoardId string       `json:"boardId" validate:"required,mongodb"`
	Status  ColumnStatus `json:"-"`
}

func (cc *ColumnCreate) ToColumnInsert() (*ColumnInsert, error) {
	boardOid, err := primitive.ObjectIDFromHex(cc.BoardId)
	if err != nil {
		return nil, common.NewBadRequestErr(errors.New("invalid objectId"))
	}

	return &ColumnInsert{
		Status:         cc.Status,
		Title:          cc.Title,
		BoardId:        boardOid,
		OrderedCardIds: []primitive.ObjectID{},
	}, nil
}
