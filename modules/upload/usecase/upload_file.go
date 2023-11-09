package uploaduc

import (
	"context"
	"pro-magnet/common"
	"pro-magnet/modules/upload/models"
)

func (uc *uploadUseCase) UploadFile(
	ctx context.Context,
	file *models.File,
) (*models.File, error) {
	url, err := uc.uploader.Upload(ctx, file.Data, file.Folder, file.Filename+"."+file.Extension)
	if err != nil {
		return nil, common.NewServerErr(err)
	}
	file.Url = url

	return file, nil
}
