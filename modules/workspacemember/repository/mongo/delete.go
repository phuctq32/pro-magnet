package wsmemberrepo

import (
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/net/context"
	"pro-magnet/common"
	wsmembermodel "pro-magnet/modules/workspacemember/model"
)

func (repo *wsMemberRepository) Delete(
	ctx context.Context,
	workspaceId, userId string,
) error {
	wsOid, err := primitive.ObjectIDFromHex(workspaceId)
	if err != nil {
		return common.NewBadRequestErr(errors.New("invalid objectId"))
	}
	userOid, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return common.NewBadRequestErr(errors.New("invalid objectId"))
	}

	if _, err = repo.db.
		Collection(wsmembermodel.WorkspaceMemberCollectionName).
		DeleteOne(ctx, bson.M{
			"workspaceId": wsOid,
			"userId":      userOid,
		}); err != nil {
		return common.NewServerErr(err)
	}

	return nil
}
