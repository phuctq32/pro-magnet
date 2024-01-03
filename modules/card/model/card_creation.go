package cardmodel

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	cardchecklistmodel "pro-magnet/modules/cardchecklist/model"
	"pro-magnet/modules/cardcomment/model"
	"time"
)

type CardCreation struct {
	ColumnId string `json:"columnId" validate:"required,mongodb"`
	Title    string `json:"title" validate:"required"`
	BoardId  string
}

type CardInsert struct {
	Id          *primitive.ObjectID                `bson:"_id,omitempty"`
	CreatedAt   time.Time                          `bson:"createdAt"`
	UpdatedAt   time.Time                          `bson:"updatedAt"`
	Status      CardStatus                         `bson:"status"`
	ColumnId    primitive.ObjectID                 `bson:"columnId"`
	BoardId     primitive.ObjectID                 `bson:"boardId"`
	Title       string                             `bson:"title"`
	Description string                             `bson:"description"`
	Cover       *string                            `bson:"cover,omitempty"`
	MemberIds   []primitive.ObjectID               `bson:"memberIds"`
	Checklists  []cardchecklistmodel.CardChecklist `bson:"checklists"`
	Comments    []cardcommentmodel.CardComment     `bson:"comments"`
	StartDate   *time.Time                         `bson:"startDate,omitempty"`
	EndDate     *time.Time                         `bson:"endDate,omitempty"`
	IsDone      bool                               `bson:"isDone"`
	Skills      []string                           `bson:"skills,omitempty"`
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
		Status:      Active,
		ColumnId:    columnOid,
		BoardId:     boardOid,
		Title:       cc.Title,
		Description: "",
		Cover:       nil,
		MemberIds:   []primitive.ObjectID{},
		Checklists:  []cardchecklistmodel.CardChecklist{},
		Comments:    []cardcommentmodel.CardComment{},
		StartDate:   nil,
		EndDate:     nil,
		IsDone:      false,
	}, nil
}
