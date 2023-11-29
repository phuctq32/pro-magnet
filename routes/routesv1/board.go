package routesv1

import (
	"github.com/gin-gonic/gin"
	"pro-magnet/components/appcontext"
	"pro-magnet/middlewares"
	boardrepo "pro-magnet/modules/board/repository/mongo"
	boardapi "pro-magnet/modules/board/transport/api"
	boarduc "pro-magnet/modules/board/usecase"
	boardmemberrepo "pro-magnet/modules/boardmember/repository/mongo"
	bmapi "pro-magnet/modules/boardmember/transport/api"
	bmuc "pro-magnet/modules/boardmember/usecase"
	wsrepo "pro-magnet/modules/workspace/repository/mongo"
)

func NewBoardRouter(appCtx appcontext.AppContext, router *gin.RouterGroup) {
	boardRepo := boardrepo.NewBoardRepository(appCtx.DBConnection())
	wsRepo := wsrepo.NewWorkspaceRepository(appCtx.DBConnection())

	boardUC := boarduc.NewBoardUseCase(boardRepo, wsRepo, appCtx.AsyncGroup())

	boardHdl := boardapi.NewBoardHandler(boardUC)

	boardRouter := router.Group("/boards", middlewares.Authorize(appCtx))
	{
		boardRouter.POST("", boardHdl.CreateBoard(appCtx))
	}

	bmRepo := boardmemberrepo.NewBoardMemberRepository(appCtx.DBConnection())
	bmUC := bmuc.NewBoardMemberUseCase(bmRepo)
	bmHdl := bmapi.NewBoardMemberHandler(bmUC)

	boardMemberRouter := boardRouter.Group("/:boardId/members")
	{
		boardMemberRouter.PATCH("", bmHdl.AddMember(appCtx))
	}
}
