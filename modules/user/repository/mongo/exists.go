package userrepo

import (
	"context"
	"pro-magnet/common"
	usermodel "pro-magnet/modules/user/model"
)

func (repo *userRepository) Exist(
	ctx context.Context,
	filter map[string]interface{},
) (bool, error) {
	count, err := repo.db.
		Collection(usermodel.UserCollectionName).
		CountDocuments(ctx, filter)
	if err != nil {
		return false, common.NewServerErr(err)
	}

	return count > 0, nil
}

func (repo *userRepository) UserExist(
	ctx context.Context,
	email string,
) (bool, error) {
	return repo.Exist(ctx, map[string]interface{}{"email": email})
}

func (repo *userRepository) InternalUserExist(
	ctx context.Context,
	email string,
) (bool, error) {
	return repo.Exist(ctx, map[string]interface{}{"email": email, "isInternal": true})
}
