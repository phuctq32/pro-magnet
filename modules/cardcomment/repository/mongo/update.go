package cardcommentrepo

import (
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/net/context"
	"pro-magnet/common"
	cardcommentmodel "pro-magnet/modules/cardcomment/model"
	"time"
)

func (repo *cardCommentRepository) Update(
	ctx context.Context,
	cardId, commentId string,
	updateData *cardcommentmodel.CardCommentUpdate,
) error {
	cardOid, err := primitive.ObjectIDFromHex(cardId)
	if err != nil {
		return common.NewBadRequestErr(errors.New("invalid objectId"))
	}
	commentOid, err := primitive.ObjectIDFromHex(commentId)
	if err != nil {
		return common.NewBadRequestErr(errors.New("invalid objectId"))
	}

	// Card filter
	filter := bson.M{"_id": cardOid}

	updateData.UpdatedAt = time.Now()
	update := bson.A{bson.M{"$set": bson.M{"comments": bson.M{
		"$map": bson.M{
			"input": "$comments",
			"in": bson.M{
				"$cond": bson.M{
					"if": bson.M{
						"$eq": bson.A{"$$this._id", commentOid},
					},
					"then": bson.M{
						"$mergeObjects": bson.A{"$$this", updateData},
					},
					"else": "$$this",
				},
			},
		},
	}}}}

	if _, err = repo.db.Collection(cardcommentmodel.CardCollectionName).
		UpdateOne(ctx, filter, update); err != nil {
		return common.NewServerErr(err)
	}

	return nil
}
