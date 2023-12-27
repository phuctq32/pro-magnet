package wsrepo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (repo *wsRepository) GetMemberIds(ctx context.Context, workspaceId string) ([]string, error) {
	wsOid, _ := primitive.ObjectIDFromHex(workspaceId)

	_, err := repo.FindOne(ctx, map[string]interface{}{"_id": wsOid})
	if err != nil {
		return nil, err
	}

	return []string{}, nil
}
