package uploadapi

import (
	"context"
	"pro-magnet/modules/upload/models"
)

type UploadUseCase interface {
	UploadFile(context.Context, *models.File) (*models.File, error)
}

type uploadHandler struct {
	uc UploadUseCase
}

func NewUploadHandler() *uploadHandler {
	return &uploadHandler{}
}
