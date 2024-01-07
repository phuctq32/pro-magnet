package boardmodel

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	columnmodel "pro-magnet/modules/column/model"
	labelmodel "pro-magnet/modules/label/model"
	usermodel "pro-magnet/modules/user/model"
	"time"
)

type BoardInsert struct {
	CreatedAt        time.Time            `bson:"createdAt"`
	UpdatedAt        time.Time            `bson:"updatedAt"`
	Status           BoardStatus          `bson:"status"`
	Name             string               `bson:"name"`
	WorkspaceId      primitive.ObjectID   `bson:"workspaceId"`
	AdminId          primitive.ObjectID   `bson:"adminId"`
	OrderedColumnIds []primitive.ObjectID `bson:"orderedColumnIds"`
}

type Board struct {
	Id               *string     `json:"_id,omitempty" bson:"_id,omitempty"`
	CreatedAt        time.Time   `json:"createdAt" bson:"createdAt"`
	UpdatedAt        time.Time   `json:"updatedAt" bson:"updatedAt"`
	Status           BoardStatus `json:"-" bson:"status"`
	Name             string      `json:"name" bson:"name"`
	WorkspaceId      string      `json:"workspaceId" bson:"workspaceId"`
	AdminId          string      `json:"adminId" bson:"adminId"`
	OrderedColumnIds []string    `json:"orderedColumnIds" bson:"orderedColumnIds"`

	// Aggregated data
	Labels           []labelmodel.Label   `json:"labels,omitempty" bson:"-"`
	FilteredLabelIds []string             `json:"filteredLabelIds,omitempty" bson:"-"`
	Columns          []columnmodel.Column `json:"columns,omitempty" bson:"-"`
	Members          []usermodel.User     `json:"members,omitempty" bson:"-"`
}
