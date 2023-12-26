package routesv1

import (
	"github.com/gin-gonic/gin"
	"pro-magnet/components/appcontext"
	"pro-magnet/middlewares"
	"pro-magnet/modules/workspace/repository/mongo"
	"pro-magnet/modules/workspace/transport/api"
	wsuc "pro-magnet/modules/workspace/usecase"
	wsmemberrepo "pro-magnet/modules/workspacemember/repository/mongo"
	wsmemberapi "pro-magnet/modules/workspacemember/transport/api"
	wsmemberuc "pro-magnet/modules/workspacemember/usecase"
)

func NewWorkspaceRouter(appCtx appcontext.AppContext, router *gin.RouterGroup) {
	wsRepo := wsrepo.NewWorkspaceRepository(appCtx.DBConnection())
	wsUC := wsuc.NewWorkspaceUseCase(wsRepo)
	wsHdl := wsapi.NewWorkspaceHandler(wsUC)

	wsRouter := router.Group("/workspaces", middlewares.Authorize(appCtx))
	{
		wsRouter.POST("", wsHdl.CreateWorkspace(appCtx))
	}

	wsMemberRepo := wsmemberrepo.NewWorkspaceMemberRepository(appCtx.DBConnection())
	wsMemberUC := wsmemberuc.NewWorkspaceMemberUseCase(wsMemberRepo, wsRepo)
	wsMemberHdl := wsmemberapi.NewWorkspaceMemberHandler(wsMemberUC)
	wsMemberRouter := wsRouter.Group("/:workspaceId/members")
	{
		wsMemberRouter.POST("", wsMemberHdl.AddWorkspaceMembers(appCtx))
		wsMemberRouter.DELETE("/:memberId", wsMemberHdl.RemoveMemberFromWorkspace(appCtx))
	}
}
