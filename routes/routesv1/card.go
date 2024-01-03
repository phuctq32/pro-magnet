package routesv1

import (
	"github.com/gin-gonic/gin"
	"pro-magnet/components/appcontext"
	"pro-magnet/middlewares"
	boardmemberrepo "pro-magnet/modules/boardmember/repository/mongo"
	"pro-magnet/modules/card/repository/mongo"
	cardapi "pro-magnet/modules/card/transport/api"
	carduc "pro-magnet/modules/card/usecase"
	carepo "pro-magnet/modules/cardattachment/repository/mongo"
	columnrepo "pro-magnet/modules/column/repository/mongo"
	labelrepo "pro-magnet/modules/label/repository/mongo"
	userrepo "pro-magnet/modules/user/repository/mongo"
)

func NewCardRouter(appCtx appcontext.AppContext, router *gin.RouterGroup) {
	cardRepo := mongo.NewCardRepository(appCtx.DBConnection())
	userRepo := userrepo.NewUserRepository(appCtx.DBConnection())
	colRepo := columnrepo.NewColumnRepository(appCtx.DBConnection())
	bmRepo := boardmemberrepo.NewBoardMemberRepository(appCtx.DBConnection())
	caRepo := carepo.NewCardAttachmentRepository(appCtx.DBConnection())
	labelRepo := labelrepo.NewLabelRepository(appCtx.DBConnection())
	cardDataAggregator := carduc.NewCardDataAggregator(appCtx.AsyncGroup(), caRepo, userRepo, labelRepo)
	cardUC := carduc.NewCardUseCase(cardRepo, colRepo, bmRepo, labelRepo, cardDataAggregator)
	cardHdl := cardapi.NewCardHandler(cardUC)

	cardRouter := router.Group("/cards", middlewares.Authorize(appCtx))
	{
		cardRouter.GET("/:cardId", cardHdl.GetCardById(appCtx))
		cardRouter.POST("", cardHdl.CreateCard(appCtx))
		cardRouter.PATCH("/:cardId", cardHdl.UpdateCardById(appCtx))
		cardRouter.DELETE("/:cardId", cardHdl.RemoveCard(appCtx))

		cardRouter.PATCH("/:cardId/date", cardHdl.UpdateCardDate(appCtx))
		cardRouter.DELETE("/:cardId/date", cardHdl.RemoveCardDate(appCtx))

		cardRouter.POST("/:cardId/members", cardHdl.AddMemberToCard(appCtx))
		cardRouter.DELETE("/:cardId/members/:memberId", cardHdl.RemoveMemberFromCard(appCtx))

		cardRouter.POST("/:cardId/labels", cardHdl.AddLabelToCard(appCtx))

		cardRouter.POST("/:cardId/skills", cardHdl.AddSkills(appCtx))
		cardRouter.DELETE("/:cardId/skills", cardHdl.RemoveSkill(appCtx))
	}
}
