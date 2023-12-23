package boardmemberrepo

import (
	"golang.org/x/net/context"
	"pro-magnet/common"
	bmmodel "pro-magnet/modules/boardmember/model"
)

func (repo *boardMemberRepository) Delete(
	ctx context.Context,
	data *bmmodel.BoardMember,
) error {
	filter, err := data.ToBoardMemberInsert()
	if err != nil {
		return err
	}
	return common.WithMongodbTransaction(ctx, repo.db, func(txCtx context.Context) error {
		if _, err := repo.db.
			Collection(bmmodel.BoardMemberCollectionName).
			DeleteOne(ctx, filter); err != nil {
			return err
		}

		return nil
	})
}
