package uploaduc

import "pro-magnet/components/upload"

type uploadUseCase struct {
	uploader upload.Uploader
}

func NewUploadUseCase(
	uploader upload.Uploader,
) *uploadUseCase {
	return &uploadUseCase{
		uploader: uploader,
	}
}
