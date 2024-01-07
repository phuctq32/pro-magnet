package cardmodel

import (
	camodel "pro-magnet/modules/cardattachment/model"
	cardchecklistmodel "pro-magnet/modules/cardchecklist/model"
	cardcommentmodel "pro-magnet/modules/cardcomment/model"
	labelmodel "pro-magnet/modules/label/model"
	usermodel "pro-magnet/modules/user/model"
	"time"
)

type CardStatus uint8

type Card struct {
	Id          *string                            `json:"_id,omitempty" bson:"_id,omitempty"`
	CreatedAt   *time.Time                         `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt   *time.Time                         `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
	Status      CardStatus                         `json:"-" bson:"status"`
	ColumnId    *string                            `json:"columnId,omitempty" bson:"columnId,omitempty"`
	BoardId     *string                            `json:"boardId,omitempty" bson:"boardId,omitempty"`
	Title       string                             `json:"title" bson:"title"`
	Description *string                            `json:"description,omitempty" bson:"description,omitempty"`
	Cover       *string                            `json:"cover" bson:"cover"`
	LabelIds    []string                           `json:"-" bson:"labelIds"`
	MemberIds   []string                           `json:"-" bson:"memberIds"`
	Checklists  []cardchecklistmodel.CardChecklist `json:"checklists,omitempty" bson:"checklists,omitempty"`
	Comments    []cardcommentmodel.CardComment     `json:"comments,omitempty" bson:"comments,omitempty"`
	StartDate   *time.Time                         `json:"startDate,omitempty" bson:"startDate,omitempty"`
	EndDate     *time.Time                         `json:"endDate,omitempty" bson:"endDate,omitempty"`
	IsDone      bool                               `json:"isDone" bson:"isDone"`
	Skills      []string                           `json:"skills,omitempty" bson:"skills,omitempty"`

	// Aggregated data
	Members         []usermodel.User         `json:"members,omitempty"`
	Labels          []labelmodel.Label       `json:"labels,omitempty"`
	Attachments     []camodel.CardAttachment `json:"attachments,omitempty"`
	MemberCount     *int                     `json:"memberCount,omitempty" bson:"memberCount,omitempty"`
	CommentCount    *int                     `json:"commentCount,omitempty" bson:"commentCount,omitempty"`
	AttachmentCount *int                     `json:"attachmentCount,omitempty"`
}
