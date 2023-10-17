package routesv1

import (
	"github.com/gin-gonic/gin"
	"pro-magnet/components/appcontext"
	uploadapi "pro-magnet/modules/upload/transport/api"
)

func NewUploadRouter(appCtx appcontext.AppContext, router *gin.RouterGroup) {
	uploadHdl := uploadapi.NewUploadHandler()

	uploadRouter := router.Group("/upload")
	{
		uploadRouter.POST("/s3", uploadHdl.UploadFileS3(appCtx))
	}
}
