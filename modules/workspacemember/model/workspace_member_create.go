package wsmembermodel

import (
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"pro-magnet/common"
)

type WorkspaceMembersCreate struct {
	UserIds     []string `json:"userIds" validate:"required,min=1,dive,mongodb"`
	WorkspaceId string   `json:"workspaceId" validate:"required,mongodb"`
}

func (wm *WorkspaceMembersCreate) ToWorkspaceMembersInsert() ([]interface{}, error) {
	wsOid, err := primitive.ObjectIDFromHex(wm.WorkspaceId)
	if err != nil {
		return nil, common.NewBadRequestErr(errors.New("invalid objectId"))
	}

	res := make([]interface{}, 0)
	for i := 0; i < len(wm.UserIds); i++ {
		userOid, err := primitive.ObjectIDFromHex(wm.UserIds[i])
		if err != nil {
			return nil, common.NewBadRequestErr(errors.New("invalid objectId"))
		}
		res = append(res, bson.M{"workspaceId": wsOid, "userId": userOid})
	}

	return res, nil
}
