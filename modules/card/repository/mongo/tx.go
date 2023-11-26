package mongo

import (
	"golang.org/x/net/context"
	"pro-magnet/common"
)

func (repo *cardRepository) WithTransaction(
	ctx context.Context,
	fn func(context.Context) error,
) error {
	return common.WithMongodbTransaction(ctx, repo.db, fn)
}
