package labelrepo

import (
	"context"
	"pro-magnet/common"
)

func (repo *labelRepository) WithTransaction(
	ctx context.Context,
	fn func(context.Context) error,
) error {
	return common.WithMongodbTransaction(ctx, repo.db, fn)
}
