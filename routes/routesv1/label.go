package routesv1

import (
	"github.com/gin-gonic/gin"
	"pro-magnet/components/appcontext"
	"pro-magnet/middlewares"
	labelrepo "pro-magnet/modules/label/repository"
	labelapi "pro-magnet/modules/label/transport/api"
	labeluc "pro-magnet/modules/label/usecase"
)

func NewLabelRouter(appCtx appcontext.AppContext, router *gin.RouterGroup) {
	labelRepo := labelrepo.NewLabelRepository(appCtx.DBConnection())

	labelUC := labeluc.NewLabelUseCase(labelRepo)

	labelHdl := labelapi.NewLabelHandler(labelUC)

	labelRouter := router.Group("/labels", middlewares.Authorize(appCtx))
	{
		labelRouter.POST("", labelHdl.CreateLabel(appCtx))
	}
}
