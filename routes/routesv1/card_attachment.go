package routesv1

import (
	"github.com/gin-gonic/gin"
	"pro-magnet/components/appcontext"
	"pro-magnet/middlewares"
	boardmemberrepo "pro-magnet/modules/boardmember/repository/mongo"
	"pro-magnet/modules/card/repository/mongo"
	carepo "pro-magnet/modules/cardattachment/repository/mongo"
	caapi "pro-magnet/modules/cardattachment/transport/api"
	cauc "pro-magnet/modules/cardattachment/usecase"
)

func NewCardAttachmentRouter(appCtx appcontext.AppContext, router *gin.RouterGroup) {
	caRepo := carepo.NewCardAttachmentRepository(appCtx.DBConnection())
	cardRepo := mongo.NewCardRepository(appCtx.DBConnection())
	bmRepo := boardmemberrepo.NewBoardMemberRepository(appCtx.DBConnection())
	caUC := cauc.NewCardAttachmentUseCase(caRepo, cardRepo, bmRepo)
	caHdl := caapi.NewCardAttachmentHandler(caUC)

	caRouter := router.Group("/cards/:cardId/attachments", middlewares.Authorize(appCtx))
	{
		caRouter.POST("", caHdl.AddCardAttachment(appCtx))
		caRouter.DELETE("/:attachmentId", caHdl.RemoveCardAttachment(appCtx))
	}
}
