package cardchecklistrepo

import (
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/net/context"
	"pro-magnet/common"
	cardchecklistmodel "pro-magnet/modules/cardchecklist/model"
)

func (repo *cardChecklistRepository) Update(
	ctx context.Context,
	cardId, checklistId string,
	data *cardchecklistmodel.CardChecklistUpdate,
) error {
	cardOid, err := primitive.ObjectIDFromHex(cardId)
	if err != nil {
		return common.NewBadRequestErr(errors.New("invalid objectId"))
	}
	checklistOid, err := primitive.ObjectIDFromHex(checklistId)
	if err != nil {
		return common.NewBadRequestErr(errors.New("invalid objectId"))
	}

	// Card filter
	filter := bson.M{"_id": cardOid}

	update := bson.A{bson.M{"$set": bson.M{"checklists": bson.M{
		"$map": bson.M{
			"input": "$checklists",
			"in": bson.M{
				"$cond": bson.M{
					"if": bson.M{
						"$eq": bson.A{"$$this._id", checklistOid},
					},
					"then": bson.M{
						"$mergeObjects": bson.A{"$$this", data},
					},
					"else": "$$this",
				},
			},
		},
	}}}}

	if _, err = repo.db.Collection(cardchecklistmodel.CardCollectionName).
		UpdateOne(ctx, filter, update); err != nil {
		return common.NewServerErr(err)
	}

	return nil
}

func (repo *cardChecklistRepository) UpdateChecklistItem(
	ctx context.Context,
	cardId, checklistId, itemId string,
	updateData *cardchecklistmodel.ChecklistItemUpdate,
) error {
	cardOid, err := primitive.ObjectIDFromHex(cardId)
	if err != nil {
		return common.NewBadRequestErr(errors.New("invalid objectId"))
	}
	checklistOid, err := primitive.ObjectIDFromHex(checklistId)
	if err != nil {
		return common.NewBadRequestErr(errors.New("invalid objectId"))
	}
	itemOid, err := primitive.ObjectIDFromHex(itemId)
	if err != nil {
		return common.NewBadRequestErr(errors.New("invalid objectId"))
	}

	filter := bson.M{"_id": cardOid}

	update := bson.A{bson.M{"$set": bson.M{"checklists": bson.M{
		"$map": bson.M{
			"input": "$checklists",
			"as":    "checklist",
			"in": bson.M{
				"$cond": bson.M{
					"if": bson.M{
						"$eq": bson.A{"$$checklist._id", checklistOid},
					},
					"then": bson.M{
						"$mergeObjects": bson.A{
							"$$checklist",
							bson.M{"items": bson.M{
								"$map": bson.M{
									"input": "$$checklist.items",
									"as":    "item",
									"in": bson.M{
										"$cond": bson.M{
											"if": bson.M{
												"$eq": bson.A{"$$item._id", itemOid},
											},
											"then": bson.M{
												"$mergeObjects": bson.A{"$$item", updateData},
											},
											"else": "$$item",
										},
									},
								},
							}},
						},
					},
					"else": "$$checklist",
				},
			},
		},
	}}}}

	if _, err = repo.db.Collection(cardchecklistmodel.CardCollectionName).
		UpdateOne(ctx, filter, update); err != nil {
		return common.NewServerErr(err)
	}

	return nil
}
