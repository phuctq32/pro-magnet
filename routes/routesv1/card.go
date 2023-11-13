package routesv1

import (
	"github.com/gin-gonic/gin"
	"pro-magnet/components/appcontext"
	"pro-magnet/middlewares"
	"pro-magnet/modules/card/repository/mongo"
	cardapi "pro-magnet/modules/card/transport/api"
	carduc "pro-magnet/modules/card/usecase"
)

func NewCardRouter(appCtx appcontext.AppContext, router *gin.RouterGroup) {
	cardRepo := mongo.NewCardRepository(appCtx.DBConnection())
	cardDataAggregator := carduc.NewCardDataAggregator(appCtx.AsyncGroup())
	cardUC := carduc.NewCardUseCase(cardRepo, cardDataAggregator)
	cardHdl := cardapi.NewCardHandler(cardUC)

	cardRouter := router.Group("/cards", middlewares.Authorize(appCtx))
	{
		cardRouter.GET("/:id", cardHdl.GetCardById(appCtx))
		cardRouter.POST("", cardHdl.CreateCard(appCtx))
		cardRouter.PATCH("/:id", cardHdl.UpdateCardById(appCtx))
	}
}
