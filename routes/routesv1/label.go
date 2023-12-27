package routesv1

import (
	"github.com/gin-gonic/gin"
	"pro-magnet/components/appcontext"
	"pro-magnet/middlewares"
	boardmemberrepo "pro-magnet/modules/boardmember/repository/mongo"
	"pro-magnet/modules/card/repository/mongo"
	labelrepo "pro-magnet/modules/label/repository/mongo"
	labelapi "pro-magnet/modules/label/transport/api"
	labeluc "pro-magnet/modules/label/usecase"
)

func NewLabelRouter(appCtx appcontext.AppContext, router *gin.RouterGroup) {
	labelRepo := labelrepo.NewLabelRepository(appCtx.DBConnection())
	cardRepo := mongo.NewCardRepository(appCtx.DBConnection())
	bmRepo := boardmemberrepo.NewBoardMemberRepository(appCtx.DBConnection())

	labelUC := labeluc.NewLabelUseCase(labelRepo, cardRepo, bmRepo)

	labelHdl := labelapi.NewLabelHandler(labelUC)

	labelRouter := router.Group("/labels", middlewares.Authorize(appCtx))
	{
		labelRouter.POST("", labelHdl.CreateLabel(appCtx))
		labelRouter.PATCH("/:labelId", labelHdl.UpdateLabel(appCtx))
		labelRouter.DELETE("/:labelId")
	}
}
