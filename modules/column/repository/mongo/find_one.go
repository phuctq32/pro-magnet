package columnrepo

import (
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
	"pro-magnet/common"
	columnmodel "pro-magnet/modules/column/model"
)

func (repo *columnRepository) FindById(ctx context.Context, id string) (*columnmodel.Column, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, common.NewBadRequestErr(errors.New("invalid objectId"))
	}

	var column columnmodel.Column
	if err := repo.db.
		Collection(columnmodel.ColumnCollectionName).
		FindOne(ctx, bson.M{"_id": oid}).Decode(&column); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, common.NewNotFoundErr("column", columnmodel.ErrColumnNotFound)
		}

		return nil, common.NewServerErr(err)
	}

	return &column, nil
}
