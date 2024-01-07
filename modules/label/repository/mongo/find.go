package labelrepo

import (
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/net/context"
	"pro-magnet/common"
	labelmodel "pro-magnet/modules/label/model"
)

func (repo *labelRepository) Find(
	ctx context.Context,
	status labelmodel.LabelStatus,
	filter map[string]interface{},
) ([]labelmodel.Label, error) {
	filter["status"] = status
	cursor, err := repo.db.Collection(labelmodel.LabelCollectionName).Find(ctx, filter)
	if err != nil {
		return nil, common.NewServerErr(err)
	}

	var labels []labelmodel.Label
	if err = cursor.All(ctx, &labels); err != nil {
		return nil, err
	}
	if labels == nil {
		return []labelmodel.Label{}, nil
	}

	return labels, nil
}

func (repo *labelRepository) FindByIds(
	ctx context.Context,
	status labelmodel.LabelStatus,
	labelIds []string,
) ([]labelmodel.Label, error) {
	labelOids := make([]primitive.ObjectID, 0)
	for _, labelId := range labelIds {
		labelOid, err := primitive.ObjectIDFromHex(labelId)
		if err != nil {
			return nil, common.NewBadRequestErr(errors.New("invalid objectId"))
		}
		labelOids = append(labelOids, labelOid)
	}

	return repo.Find(ctx, status, map[string]interface{}{"_id": bson.M{"$in": labelOids}})
}

func (repo *labelRepository) FindByBoardId(
	ctx context.Context,
	labelStatus labelmodel.LabelStatus,
	boardId string,
) ([]labelmodel.Label, error) {
	boardOid, err := primitive.ObjectIDFromHex(boardId)
	if err != nil {
		return nil, common.NewBadRequestErr(errors.New("invalid objectId"))
	}

	filter := map[string]interface{}{"boardId": boardOid}

	return repo.Find(ctx, labelStatus, filter)
}
