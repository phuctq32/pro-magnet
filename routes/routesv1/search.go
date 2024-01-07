package routesv1

import (
	"github.com/gin-gonic/gin"
	"pro-magnet/components/appcontext"
	"pro-magnet/middlewares"
	boardrepo "pro-magnet/modules/board/repository/mongo"
	"pro-magnet/modules/card/repository/mongo"
	searchapi "pro-magnet/modules/search/transport/api"
	searchuc "pro-magnet/modules/search/usecase"
	wsrepo "pro-magnet/modules/workspace/repository/mongo"
)

func NewSearchRouter(appCtx appcontext.AppContext, router *gin.RouterGroup) {
	wsRepo := wsrepo.NewWorkspaceRepository(appCtx.DBConnection())
	boardRepo := boardrepo.NewBoardRepository(appCtx.DBConnection())
	cardRepo := mongo.NewCardRepository(appCtx.DBConnection())
	searchUC := searchuc.NewSearchUseCase(wsRepo, boardRepo, cardRepo, appCtx.AsyncGroup())

	searchHdl := searchapi.NewSearchHandler(searchUC)

	searchRouter := router.Group("/search", middlewares.Authorize(appCtx))
	{
		searchRouter.GET("", searchHdl.Search(appCtx))
	}
}
