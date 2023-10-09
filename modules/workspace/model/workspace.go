package wsmodel

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const (
	WsCollectionName = "workspaces"
	DefaultImageUrl  = "https://img.fruugo.com/product/2/87/557318872_max.jpg"
)

type WorkspaceInsert struct {
	CreatedAt time.Time          `bson:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt"`
	UserId    primitive.ObjectID `bson:"userId"`
	Name      string             `bson:"name"`
	Image     string             `bson:"image"`
}

type Workspace struct {
	Id        *string   `json:"_id,omitempty" bson:"_id,omitempty"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"`
	UserId    string    `json:"userId" bson:"userId"`
	Name      string    `json:"name" bson:"name"`
	Image     string    `json:"image" bson:"image"`
}
