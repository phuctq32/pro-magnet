package uploadapi

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"pro-magnet/common"
	"pro-magnet/components/appcontext"
	"time"
)

func (hdl *uploadHandler) UploadFileWithCloudinary(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		file, err := hdl.getFileFromRequest(c)
		if err != nil {
			panic(err)
		}
		file.Folder = fmt.Sprintf("%v", time.Now().UnixNano())

		data, err := hdl.cldUploadUC.UploadFile(c.Request.Context(), file)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewResponse("uploaded file", data))
	}
}
