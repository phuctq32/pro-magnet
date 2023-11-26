package columnrepo

import (
	"golang.org/x/net/context"
	"pro-magnet/common"
)

func (repo *columnRepository) WithTransaction(
	ctx context.Context,
	fn func(context.Context) error,
) error {
	return common.WithMongodbTransaction(ctx, repo.db, fn)
}
