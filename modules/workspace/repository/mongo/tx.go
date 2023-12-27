package wsrepo

import (
	"golang.org/x/net/context"
	"pro-magnet/common"
)

func (repo *wsRepository) WithTransaction(
	ctx context.Context,
	fn func(context.Context) error,
) error {
	return common.WithMongodbTransaction(ctx, repo.db, fn)
}
