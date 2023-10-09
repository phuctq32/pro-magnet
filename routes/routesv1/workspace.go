package routesv1

import (
	"github.com/gin-gonic/gin"
	"pro-magnet/components/appcontext"
	"pro-magnet/middlewares"
	"pro-magnet/modules/workspace/repository"
	"pro-magnet/modules/workspace/transport/api"
	wsuc "pro-magnet/modules/workspace/usecase"
)

func NewWorkspaceRouter(appCtx appcontext.AppContext, router *gin.RouterGroup) {
	wsRepo := wsrepo.NewWorkspaceRepository(appCtx.DBConnection())
	wsUC := wsuc.NewWorkspaceUseCase(wsRepo)
	wsHdl := wsapi.NewWorkspaceHandler(wsUC)

	wsRouter := router.Group("/workspaces", middlewares.Authorize(appCtx))
	{
		wsRouter.POST("", wsHdl.CreateWorkspace(appCtx))
	}
}
