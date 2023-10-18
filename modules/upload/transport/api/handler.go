package uploadapi

import (
	"context"
	"pro-magnet/modules/upload/models"
)

type UploadUseCase interface {
	UploadFile(context.Context, *models.File) (*models.File, error)
}

type uploadHandler struct {
	s3UploadUC  UploadUseCase
	cldUploadUC UploadUseCase
}

func NewUploadHandler(
	s3UploadUC UploadUseCase,
	cldUploadUC UploadUseCase,
) *uploadHandler {
	return &uploadHandler{
		s3UploadUC:  s3UploadUC,
		cldUploadUC: cldUploadUC,
	}
}
