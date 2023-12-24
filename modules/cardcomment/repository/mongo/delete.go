package cardcommentrepo

import (
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/net/context"
	"pro-magnet/common"
	cardcommentmodel "pro-magnet/modules/cardcomment/model"
)

func (repo *cardCommentRepository) Delete(
	ctx context.Context,
	cardId, commentId string,
) error {
	cardOid, err := primitive.ObjectIDFromHex(cardId)
	if err != nil {
		return common.NewBadRequestErr(errors.New("invalid objectId"))
	}
	commentOid, err := primitive.ObjectIDFromHex(commentId)
	if err != nil {
		return common.NewBadRequestErr(errors.New("invalid objectId"))
	}

	filter := bson.M{"_id": cardOid}
	update := bson.M{"$pull": bson.M{"comments": bson.M{"_id": commentOid}}}

	if _, err := repo.db.Collection(cardcommentmodel.CardCollectionName).
		UpdateOne(ctx, filter, update); err != nil {
		return common.NewServerErr(err)
	}

	return nil
}
