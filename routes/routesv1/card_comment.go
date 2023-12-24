package routesv1

import (
	"github.com/gin-gonic/gin"
	"pro-magnet/components/appcontext"
	"pro-magnet/middlewares"
	boardmemberrepo "pro-magnet/modules/boardmember/repository/mongo"
	"pro-magnet/modules/card/repository/mongo"
	cardcommentrepo "pro-magnet/modules/cardcomment/repository/mongo"
	cardcommentapi "pro-magnet/modules/cardcomment/transport/api"
	cardcommentuc "pro-magnet/modules/cardcomment/usecase"
)

func NewCardCommentRouter(appCtx appcontext.AppContext, router *gin.RouterGroup) {
	cardRepo := mongo.NewCardRepository(appCtx.DBConnection())
	bmRepo := boardmemberrepo.NewBoardMemberRepository(appCtx.DBConnection())
	cmRepo := cardcommentrepo.NewCardCommentRepository(appCtx.DBConnection())

	cmUC := cardcommentuc.NewCardCommentUseCase(cmRepo, cardRepo, bmRepo)

	cmHdl := cardcommentapi.NewCardCommentHandler(cmUC)

	cmRouter := router.Group("/cards/:cardId/comments", middlewares.Authorize(appCtx))
	{
		cmRouter.POST("", cmHdl.CreateCardComment(appCtx))
		cmRouter.PATCH("/:commentId", cmHdl.UpdateCardComment(appCtx))
		cmRouter.DELETE("/:commentId", cmHdl.DeleteComment(appCtx))
	}
}
