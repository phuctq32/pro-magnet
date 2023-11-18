package routesv1

import (
	"github.com/gin-gonic/gin"
	"pro-magnet/components/appcontext"
	"pro-magnet/middlewares"
	"pro-magnet/modules/card/repository/mongo"
	cardchecklistrepo "pro-magnet/modules/cardchecklist/repository/mongo"
	cardchecklistapi "pro-magnet/modules/cardchecklist/transport/api"
	cardchecklistuc "pro-magnet/modules/cardchecklist/usecase"
)

func NewCardChecklistRouter(appCtx appcontext.AppContext, router *gin.RouterGroup) {
	ccRepo := cardchecklistrepo.NewCardChecklistRepository(appCtx.DBConnection())
	cardRepo := mongo.NewCardRepository(appCtx.DBConnection())
	ccUC := cardchecklistuc.NewCardChecklistUseCase(ccRepo, cardRepo)
	ccHdl := cardchecklistapi.NewCardChecklistHandler(ccUC)

	ccRouter := router.Group("/cards/:cardId/checklists", middlewares.Authorize(appCtx))
	{
		ccRouter.POST("/:checklistId/items")
		ccRouter.PATCH("/:checklistId/items/:itemId")
		ccRouter.DELETE("/:checklistId/items/:itemId")

		ccRouter.POST("", ccHdl.CreateChecklist(appCtx))
		ccRouter.PATCH("/:checklistId")
		ccRouter.DELETE("/:checklistId")
	}
}
