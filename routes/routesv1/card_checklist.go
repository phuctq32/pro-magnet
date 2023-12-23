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
		ccRouter.POST("/:checklistId/items", ccHdl.CreateChecklistItem(appCtx))
		ccRouter.PATCH("/:checklistId/items/:itemId", ccHdl.UpdateChecklistItem(appCtx))
		ccRouter.DELETE("/:checklistId/items/:itemId", ccHdl.DeleteChecklistItem(appCtx))

		ccRouter.POST("", ccHdl.CreateChecklist(appCtx))
		ccRouter.PATCH("/:checklistId", ccHdl.UpdateCardChecklist(appCtx))
		ccRouter.DELETE("/:checklistId", ccHdl.DeleteChecklist(appCtx))
	}
}
