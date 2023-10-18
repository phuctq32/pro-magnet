package routesv1

import (
	"github.com/gin-gonic/gin"
	"pro-magnet/components/appcontext"
	uploadapi "pro-magnet/modules/upload/transport/api"
	uploaduc "pro-magnet/modules/upload/usecase"
)

func NewUploadRouter(appCtx appcontext.AppContext, router *gin.RouterGroup) {
	uploadHdl := uploadapi.NewUploadHandler(
		uploaduc.NewUploadUseCase(appCtx.S3Uploader()),
		uploaduc.NewUploadUseCase(appCtx.CloudinaryUploader()),
	)

	uploadRouter := router.Group("/upload")
	{
		uploadRouter.POST("/s3", uploadHdl.UploadFileWithS3(appCtx))
		uploadRouter.POST("/cloudinary", uploadHdl.UploadFileWithCloudinary(appCtx))
	}
}
