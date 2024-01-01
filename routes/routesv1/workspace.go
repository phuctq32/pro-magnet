package routesv1

import (
	"github.com/gin-gonic/gin"
	"pro-magnet/components/appcontext"
	"pro-magnet/middlewares"
	boardrepo "pro-magnet/modules/board/repository/mongo"
	userrepo "pro-magnet/modules/user/repository/mongo"
	"pro-magnet/modules/workspace/repository/mongo"
	"pro-magnet/modules/workspace/transport/api"
	wsuc "pro-magnet/modules/workspace/usecase"
	wsmemberrepo "pro-magnet/modules/workspacemember/repository/mongo"
	wsmemberapi "pro-magnet/modules/workspacemember/transport/api"
	wsmemberuc "pro-magnet/modules/workspacemember/usecase"
)

func NewWorkspaceRouter(appCtx appcontext.AppContext, router *gin.RouterGroup) {
	wsRepo := wsrepo.NewWorkspaceRepository(appCtx.DBConnection())
	wsMemberRepo := wsmemberrepo.NewWorkspaceMemberRepository(appCtx.DBConnection())
	userRepo := userrepo.NewUserRepository(appCtx.DBConnection())
	boardRepo := boardrepo.NewBoardRepository(appCtx.DBConnection())
	wsAgg := wsuc.NewWorkspaceAggregator(appCtx.AsyncGroup(), userRepo, wsMemberRepo, boardRepo)
	wsUC := wsuc.NewWorkspaceUseCase(wsRepo, wsMemberRepo, wsAgg)
	wsHdl := wsapi.NewWorkspaceHandler(wsUC)

	wsRouter := router.Group("/workspaces", middlewares.Authorize(appCtx))
	{
		wsRouter.POST("", wsHdl.CreateWorkspace(appCtx))
		wsRouter.GET("", wsHdl.GetCurrentUserWorkspaces(appCtx))
		wsRouter.GET("/:workspaceId", wsHdl.GetWorkspaceById(appCtx))
		wsRouter.PATCH("/:workspaceId", wsHdl.UpdateWorkspace(appCtx))
	}

	wsMemberUC := wsmemberuc.NewWorkspaceMemberUseCase(wsMemberRepo, wsRepo)
	wsMemberHdl := wsmemberapi.NewWorkspaceMemberHandler(wsMemberUC)
	wsMemberRouter := wsRouter.Group("/:workspaceId/members")
	{
		wsMemberRouter.POST("", wsMemberHdl.AddWorkspaceMembers(appCtx))
		wsMemberRouter.DELETE("/:memberId", wsMemberHdl.RemoveMemberFromWorkspace(appCtx))
	}
}
