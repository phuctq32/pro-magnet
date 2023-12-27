package boardmodel

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	labelmodel "pro-magnet/modules/label/model"
	"time"
)

type BoardInsert struct {
	CreatedAt      time.Time            `bson:"createdAt"`
	UpdatedAt      time.Time            `bson:"updatedAt"`
	Status         BoardStatus          `bson:"status"`
	Name           string               `bson:"name"`
	WorkspaceId    primitive.ObjectID   `bson:"workspaceId"`
	AdminId        primitive.ObjectID   `bson:"adminId"`
	ColumnOrderIds []primitive.ObjectID `bson:"columnOrderIds"`
}

type Board struct {
	Id             *string     `json:"_id,omitempty" bson:"_id,omitempty"`
	CreatedAt      time.Time   `json:"createdAt" bson:"createdAt"`
	UpdatedAt      time.Time   `json:"updatedAt" bson:"updatedAt"`
	Status         BoardStatus `json:"-" bson:"status"`
	Name           string      `json:"name" bson:"name"`
	WorkspaceId    string      `json:"workspaceId" bson:"workspaceId"`
	AdminId        string      `json:"adminId" bson:"adminId"`
	ColumnOrderIds []string    `json:"columnOrderIds" bson:"columnOrderIds"`

	// Aggregated data
	Labels []labelmodel.Label `json:"labels" bson:"-"`
}
