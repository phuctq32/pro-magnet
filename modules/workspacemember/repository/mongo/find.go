package wsmemberrepo

import (
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/net/context"
	"pro-magnet/common"
	wsmembermodel "pro-magnet/modules/workspacemember/model"
)

func (repo *wsMemberRepository) FindMemberIdsByWorkspaceId(
	ctx context.Context,
	workspaceId string,
) ([]string, error) {
	wsOid, err := primitive.ObjectIDFromHex(workspaceId)
	if err != nil {
		return nil, common.NewBadRequestErr(errors.New("invalid objectId"))
	}

	cursor, err := repo.db.
		Collection(wsmembermodel.WorkspaceMemberCollectionName).
		Find(ctx, bson.M{"workspaceId": wsOid})
	if err != nil {
		return nil, common.NewServerErr(err)
	}

	var wsMembers []wsmembermodel.WorkspaceMember
	if err = cursor.All(ctx, &wsMembers); err != nil {
		return nil, common.NewServerErr(err)
	}

	if wsMembers == nil {
		return []string{}, nil
	}

	result := make([]string, 0)
	for i := 0; i < len(wsMembers); i++ {
		result = append(result, wsMembers[i].UserId)
	}

	return result, nil
}
