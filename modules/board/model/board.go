package boardmodel

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const (
	BoardCollectionName string = "boards"
)

type BoardInsert struct {
	CreatedAt      time.Time            `bson:"createdAt"`
	UpdatedAt      time.Time            `bson:"updatedAt"`
	Name           string               `bson:"name"`
	WorkspaceId    primitive.ObjectID   `bson:"workspaceId"`
	AdminId        primitive.ObjectID   `bson:"adminId"`
	ColumnOrderIds []primitive.ObjectID `bson:"columnOrderIds"`
	//Labels         []Label   `bson:"labels"`
}

type Board struct {
	Id             *string   `json:"_id,omitempty" bson:"_id,omitempty"`
	CreatedAt      time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt" bson:"updatedAt"`
	Name           string    `json:"name" bson:"name"`
	WorkspaceId    string    `json:"workspaceId" bson:"workspaceId"`
	AdminId        string    `json:"adminId" bson:"adminId"`
	ColumnOrderIds []string  `json:"columnOrderIds" bson:"columnOrderIds"`
	//Labels         []Label   `json:"labels" bson:"labels"`
}
