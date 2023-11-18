package cardmodel

import (
	camodel "pro-magnet/modules/cardattachment/model"
	cardchecklistmodel "pro-magnet/modules/cardchecklist/model"
	labelmodel "pro-magnet/modules/label/model"
	"time"
)

type CardStatus uint8

type Card struct {
	Id          *string                            `json:"_id,omitempty" bson:"_id,omitempty"`
	CreatedAt   time.Time                          `json:"createdAt" bson:"createdAt"`
	UpdatedAt   time.Time                          `json:"updatedAt" bson:"updatedAt"`
	Status      CardStatus                         `json:"-" bson:"status"`
	ColumnId    string                             `json:"columnId" bson:"columnId"`
	BoardId     string                             `json:"boardId" bson:"boardId"`
	Title       string                             `json:"title" bson:"title"`
	Description string                             `json:"description" bson:"description"`
	Cover       *string                            `json:"cover" bson:"cover,omitempty"`
	MemberIds   []string                           `json:"-" bson:"memberIds"`
	Checklists  []cardchecklistmodel.CardChecklist `json:"checklists" bson:"checklists"`
	Comments    []CardComment                      `json:"comments" bson:"comments"`
	StartDate   *time.Time                         `json:"startDate" bson:"startDate,omitempty"`
	EndDate     *time.Time                         `json:"endDate" bson:"endDate,omitempty"`
	IsDone      bool                               `json:"isDone" bson:"isDone"`

	// Aggregated data
	Labels      []labelmodel.Label       `json:"labels"`
	Attachments []camodel.CardAttachment `json:"attachments"`
}
