package userrepo

import (
	"context"
	"pro-magnet/common"
)

// Exists return the error when user was found or an error occurred
func (repo *userRepository) Exists(ctx context.Context, filter map[string]interface{}) error {
	_, err := repo.FindOne(ctx, filter)
	if err == nil {
		return common.NewExistedErr("user")
	} else if err.Error() != common.ErrNotFound.Error() {
		return err
	}

	return nil
}

func (repo *userRepository) CheckEmailExists(ctx context.Context, email string) error {
	return repo.Exists(ctx, map[string]interface{}{"email": email})
}
