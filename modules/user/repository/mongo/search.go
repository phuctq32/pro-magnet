package userrepo

import (
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
	"pro-magnet/common"
	usermodel "pro-magnet/modules/user/model"
)

func (repo *userRepository) SearchUsersByEmail(
	ctx context.Context,
	emailSearchStr string,
	exceptedUserIds []string,
) ([]usermodel.User, error) {
	exceptedUserOids := make([]primitive.ObjectID, 0)
	for _, userId := range exceptedUserIds {
		oid, err := primitive.ObjectIDFromHex(userId)
		if err != nil {
			return nil, common.NewBadRequestErr(errors.New("invalid objectId"))
		}
		exceptedUserOids = append(exceptedUserOids, oid)
	}

	emailSearchRegex := primitive.Regex{
		Pattern: emailSearchStr,
		Options: "i",
	}
	filter := bson.M{
		"_id":        bson.M{"$nin": exceptedUserOids},
		"isVerified": true,
		//"$text":      bson.M{"$search": emailSearchStr},
		"email": bson.M{"$regex": emailSearchRegex},
	}

	opts := options.Find()
	opts.SetProjection(bson.M{
		"_id":    1,
		"name":   1,
		"email":  1,
		"avatar": 1,
	})
	opts.SetLimit(10)

	cursor, err := repo.db.
		Collection(usermodel.UserCollectionName).
		Find(ctx, filter, opts)
	if err != nil {
		return nil, common.NewServerErr(err)
	}

	var users []usermodel.User
	if err = cursor.All(ctx, &users); err != nil {
		return nil, common.NewServerErr(err)
	}

	if users == nil {
		return []usermodel.User{}, nil
	}

	return users, nil
}
