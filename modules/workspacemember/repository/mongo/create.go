package wsmemberrepo

import (
	"golang.org/x/net/context"
	"pro-magnet/common"
	wsmembermodel "pro-magnet/modules/workspacemember/model"
)

func (repo *wsMemberRepository) CreateMany(
	ctx context.Context,
	data *wsmembermodel.WorkspaceMembersCreate,
) error {
	insertData, err := data.ToWorkspaceMembersInsert()
	if err != nil {
		return err
	}

	if _, err = repo.db.
		Collection(wsmembermodel.WorkspaceMemberCollectionName).
		InsertMany(ctx, insertData); err != nil {
		return common.NewServerErr(err)
	}

	return nil
}
