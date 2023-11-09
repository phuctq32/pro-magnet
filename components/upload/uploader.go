package upload

import "context"

type Uploader interface {
	Upload(ctx context.Context, data []byte, folder, filename string) (string, error)
}
