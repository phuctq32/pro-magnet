package uploadapi

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
	"pro-magnet/common"
	"pro-magnet/components/appcontext"
	"pro-magnet/modules/upload/models"
	uploaduc "pro-magnet/modules/upload/usecase"
	"strings"
	"time"
)

func (hdl *uploadHandler) UploadFileS3(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		fileHeader, err := c.FormFile("file")

		if err != nil {
			panic(common.NewBadRequestErr(errors.New("could not get file")))
		}
		file, err := fileHeader.Open()
		if err != nil {
			panic(common.NewBadRequestErr(errors.New("could not get file")))
		}
		defer func() {
			_ = file.Close()
		}()

		dataBytes := make([]byte, fileHeader.Size)
		if _, err := file.Read(dataBytes); err != nil {
			panic(err)
		}

		hdl.uc = uploaduc.NewUploadUseCase(appCtx.S3Uploader())

		str := strings.Split(fileHeader.Filename, ".")
		data, err := hdl.uc.UploadFile(c.Request.Context(), &models.File{
			Filename:  str[0],
			Extension: str[1],
			Data:      dataBytes,
			Folder:    fmt.Sprintf("%v", time.Now().UnixNano()),
		})
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewResponse("uploaded file", data))
	}
}
