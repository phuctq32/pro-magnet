package cardcommentrepo

import (
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/net/context"
	"pro-magnet/common"
	cardmodel "pro-magnet/modules/card/model"
	cardcommentmodel "pro-magnet/modules/cardcomment/model"
	"time"
)

func (repo *cardCommentRepository) Create(
	ctx context.Context,
	cardId string,
	data *cardcommentmodel.CardCommentCreate,
) error {
	cardOid, err := primitive.ObjectIDFromHex(cardId)
	if err != nil {
		return common.NewBadRequestErr(errors.New("invalid objectId"))
	}
	authorOid, err := primitive.ObjectIDFromHex(data.UserId)
	if err != nil {
		return common.NewBadRequestErr(errors.New("invalid objectId"))
	}

	// Card filter
	filter := bson.M{"_id": cardOid}

	now := time.Now()
	insertData := &cardcommentmodel.CardCommentInsert{
		Id:        primitive.NewObjectID(),
		CreatedAt: now,
		UpdatedAt: now,
		Content:   data.Content,
		AuthorId:  authorOid,
	}
	update := bson.M{"$push": bson.M{"comments": insertData}}

	result, err := repo.db.
		Collection(cardcommentmodel.CardCollectionName).
		UpdateOne(ctx, filter, update)
	if err != nil {
		return common.NewServerErr(err)
	}
	if result.MatchedCount == 0 {
		return common.NewBadRequestErr(cardmodel.ErrCardNotFound)
	}

	return nil
}
