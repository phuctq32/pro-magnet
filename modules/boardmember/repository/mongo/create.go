package boardmemberrepo

import (
	"golang.org/x/net/context"
	"pro-magnet/common"
	bmmodel "pro-magnet/modules/boardmember/model"
)

func (repo *boardMemberRepository) Create(
	ctx context.Context,
	data *bmmodel.BoardMember,
) error {
	insertData, err := data.ToBoardMemberInsert()
	if err != nil {
		return err
	}

	if _, err = repo.db.
		Collection(bmmodel.BoardMemberCollectionName).
		InsertOne(ctx, insertData); err != nil {
		return common.NewServerErr(err)
	}

	return nil
}
