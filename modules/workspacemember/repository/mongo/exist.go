package wsmemberrepo

import (
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/net/context"
	"pro-magnet/common"
	wsmembermodel "pro-magnet/modules/workspacemember/model"
)

func (repo *wsMemberRepository) IsExist(
	ctx context.Context,
	filter map[string]interface{},
) (bool, error) {
	count, err := repo.db.
		Collection(wsmembermodel.WorkspaceMemberCollectionName).
		CountDocuments(ctx, filter)
	if err != nil {
		return false, common.NewServerErr(err)
	}

	return count > 0, nil
}

func (repo *wsMemberRepository) IsWorkspaceMember(
	ctx context.Context,
	workspaceId, userId string,
) (bool, error) {
	wsOid, err := primitive.ObjectIDFromHex(workspaceId)
	if err != nil {
		return false, common.NewBadRequestErr(errors.New("invalid object Id"))
	}
	userOid, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return false, common.NewBadRequestErr(errors.New("invalid object Id"))
	}

	return repo.IsExist(ctx, map[string]interface{}{
		"workspaceId": wsOid,
		"userId":      userOid,
	})
}
