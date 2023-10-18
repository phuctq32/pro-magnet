package upload

import (
	"bytes"
	"context"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type cldUploader struct {
	cld *cloudinary.Cloudinary
}

func NewCloudinaryUploader(cloudName, apiKey, apiSecret string) (Uploader, error) {
	cld, err := cloudinary.NewFromParams(cloudName, apiKey, apiSecret)
	if err != nil {
		return nil, err
	}
	return &cldUploader{cld: cld}, nil
}

func (cld *cldUploader) Upload(ctx context.Context, data []byte, folder, filename string) (string, error) {
	fileBytes := bytes.NewBuffer(data)

	uploadRes, err := cld.cld.Upload.Upload(ctx, fileBytes, uploader.UploadParams{
		PublicID: filename,
		Folder:   folder,
	})
	if err != nil {
		return "", err
	}

	return uploadRes.URL, err
}
