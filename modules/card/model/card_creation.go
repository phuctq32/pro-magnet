package cardmodel

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type CardCreation struct {
	ColumnId string `json:"columnId" validate:"required,mongodb"`
	Title    string `json:"title" validate:"required"`
	BoardId  string
}

type CardInsert struct {
	Id          *primitive.ObjectID  `bson:"_id,omitempty"`
	CreatedAt   time.Time            `bson:"createdAt"`
	UpdatedAt   time.Time            `bson:"updatedAt"`
	ColumnId    primitive.ObjectID   `bson:"columnId"`
	BoardId     primitive.ObjectID   `bson:"boardId"`
	Title       string               `bson:"title"`
	Description string               `bson:"description"`
	Cover       *string              `bson:"cover,omitempty"`
	MemberIds   []primitive.ObjectID `bson:"memberIds"`
	Checklists  []CardChecklist      `bson:"checklists"`
	Comments    []CardComment        `bson:"comments"`
	StartDate   *time.Time           `bson:"startDate,omitempty"`
	EndDate     *time.Time           `bson:"endDate,omitempty"`
	IsDone      bool                 `bson:"isDone"`
}

func (cc *CardCreation) ToCardInsert() (*CardInsert, error) {
	now := time.Now()

	columnOid, err := primitive.ObjectIDFromHex(cc.ColumnId)
	if err != nil {
		return nil, err
	}

	boardOid, err := primitive.ObjectIDFromHex(cc.BoardId)
	if err != nil {
		return nil, err
	}

	return &CardInsert{
		CreatedAt:   now,
		UpdatedAt:   now,
		ColumnId:    columnOid,
		BoardId:     boardOid,
		Title:       cc.Title,
		Description: "",
		Cover:       nil,
		MemberIds:   []primitive.ObjectID{},
		Checklists:  []CardChecklist{},
		Comments:    []CardComment{},
		StartDate:   nil,
		EndDate:     nil,
		IsDone:      false,
	}, nil
}
