package routesv1

import (
	"github.com/gin-gonic/gin"
	"pro-magnet/components/appcontext"
	"pro-magnet/middlewares"
	boardmemberrepo "pro-magnet/modules/boardmember/repository/mongo"
	"pro-magnet/modules/card/repository/mongo"
	columnrepo "pro-magnet/modules/column/repository/mongo"
	columnapi "pro-magnet/modules/column/transport/api"
	columnuc "pro-magnet/modules/column/usecase"
)

func NewColumnRouter(appCtx appcontext.AppContext, router *gin.RouterGroup) {
	colRepo := columnrepo.NewColumnRepository(appCtx.DBConnection())
	bmRepo := boardmemberrepo.NewBoardMemberRepository(appCtx.DBConnection())
	cardRepo := mongo.NewCardRepository(appCtx.DBConnection())
	colUC := columnuc.NewColumnUseCase(colRepo, bmRepo, cardRepo)
	colHdl := columnapi.NewColumnHandler(colUC)

	colRouter := router.Group("/columns", middlewares.Authorize(appCtx))
	{
		colRouter.POST("", colHdl.CreateColumn(appCtx))
		colRouter.PATCH("/:id", colHdl.UpdateColumn(appCtx))
		colRouter.DELETE("/:id", colHdl.RemoveColumn(appCtx))
	}
}
