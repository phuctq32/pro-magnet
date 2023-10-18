package uploadapi

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"pro-magnet/common"
	"pro-magnet/modules/upload/models"
	"strings"
)

func (hdl *uploadHandler) getFileFromRequest(c *gin.Context) (*models.File, error) {
	fileHeader, err := c.FormFile("file")

	if err != nil {
		return nil, common.NewBadRequestErr(errors.New("could not get file"))
	}
	file, err := fileHeader.Open()
	if err != nil {
		return nil, common.NewBadRequestErr(errors.New("could not get file"))
	}
	defer func() {
		_ = file.Close()
	}()

	dataBytes := make([]byte, fileHeader.Size)
	if _, err := file.Read(dataBytes); err != nil {
		return nil, err
	}

	str := strings.Split(fileHeader.Filename, ".")

	return &models.File{
		Filename:  str[0],
		Extension: str[1],
		Data:      dataBytes,
	}, nil
}
