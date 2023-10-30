package userrepo

import (
	"context"
	"pro-magnet/common"
)

func (repo *userRepository) WithTransaction(
	ctx context.Context,
	fn func(context.Context) error,
) error {
	return common.WithMongodbTransaction(ctx, repo.db, fn)
}
