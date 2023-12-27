package labelrepo

import (
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
	"pro-magnet/common"
	labelmodel "pro-magnet/modules/label/model"
)

func (repo *labelRepository) FindOne(
	ctx context.Context,
	filter map[string]interface{},
) (*labelmodel.Label, error) {
	var label labelmodel.Label
	if err := repo.db.
		Collection(labelmodel.LabelCollectionName).
		FindOne(ctx, filter).
		Decode(&label); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, common.NewNotFoundErr("label", mongo.ErrNoDocuments)
		}

		return nil, common.NewServerErr(err)
	}

	return &label, nil
}

func (repo *labelRepository) FindById(
	ctx context.Context,
	labelId string,
) (*labelmodel.Label, error) {
	labelOid, err := primitive.ObjectIDFromHex(labelId)
	if err != nil {
		return nil, err
	}

	return repo.FindOne(ctx, map[string]interface{}{"_id": labelOid})
}
