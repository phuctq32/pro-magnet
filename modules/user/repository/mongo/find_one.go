package userrepo

import (
	"context"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"pro-magnet/common"
	usermodel "pro-magnet/modules/user/model"
)

func (repo *userRepository) FindOne(ctx context.Context, filter map[string]interface{}) (*usermodel.User, error) {
	var user usermodel.User

	if err := repo.db.
		Collection(usermodel.UserCollectionName).
		FindOne(ctx, filter).
		Decode(&user); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, common.NewNotFoundErr("user", mongo.ErrNoDocuments)
		}
		return nil, common.NewServerErr(err)
	}

	return &user, nil
}

func (repo *userRepository) FindByEmail(ctx context.Context, email string) (*usermodel.User, error) {
	return repo.FindOne(ctx, map[string]interface{}{"email": email})
}
