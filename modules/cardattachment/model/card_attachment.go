package camodel

import (
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"pro-magnet/common"
	"time"
)

type CardAttachmentStatus uint8

type CardAttachmentInsert struct {
	CreatedAt time.Time            `bson:"createdAt"`
	Status    CardAttachmentStatus `bson:"status"`
	BoardId   primitive.ObjectID   `bson:"boardId"`
	CardId    primitive.ObjectID   `bson:"cardId"`
	URL       string               `bson:"url"`
	FileName  string               `bson:"fileName"`
	Extension string               `bson:"extension"`
}

type CardAttachment struct {
	Id        *string              `json:"_id,omitempty" bson:"_id,omitempty"`
	CreatedAt time.Time            `json:"createdAt" bson:"createdAt"`
	Status    CardAttachmentStatus `json:"-" bson:"status"`
	BoardId   string               `json:"-" bson:"boardId"`
	CardId    string               `json:"cardId" bson:"cardId"`
	URL       string               `json:"url" bson:"url" validate:"required,url"`
	FileName  string               `json:"fileName" bson:"fileName" validate:"required"`
	Extension string               `json:"extension" bson:"extension" validate:"required"`
}

func (ca *CardAttachment) ToDataInsert() (*CardAttachmentInsert, error) {
	cardOid, err := primitive.ObjectIDFromHex(ca.CardId)
	if err != nil {
		return nil, common.NewBadRequestErr(errors.New("invalid object id"))
	}

	boardOid, err := primitive.ObjectIDFromHex(ca.CardId)
	if err != nil {
		return nil, common.NewBadRequestErr(errors.New("invalid object id"))
	}

	return &CardAttachmentInsert{
		CreatedAt: time.Now(),
		Status:    ca.Status,
		BoardId:   boardOid,
		CardId:    cardOid,
		URL:       ca.URL,
		FileName:  ca.FileName,
		Extension: ca.Extension,
	}, nil
}
