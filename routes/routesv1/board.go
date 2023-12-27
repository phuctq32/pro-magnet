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
	userrepo "pro-magnet/modules/user/repository/mongo"
	wsmemberrepo "pro-magnet/modules/workspacemember/repository/mongo"
)

func NewBoardRouter(appCtx appcontext.AppContext, router *gin.RouterGroup) {
	boardRepo := boardrepo.NewBoardRepository(appCtx.DBConnection())
	bmRepo := boardmemberrepo.NewBoardMemberRepository(appCtx.DBConnection())
	wsMemberRepo := wsmemberrepo.NewWorkspaceMemberRepository(appCtx.DBConnection())

	boardUC := boarduc.NewBoardUseCase(boardRepo, bmRepo, wsMemberRepo, appCtx.AsyncGroup())

	boardHdl := boardapi.NewBoardHandler(boardUC)

	boardRouter := router.Group("/boards", middlewares.Authorize(appCtx))
	{
		boardRouter.POST("", boardHdl.CreateBoard(appCtx))
		boardRouter.PATCH("/:boardId", boardHdl.UpdateBoard(appCtx))
	}

	userRepo := userrepo.NewUserRepository(appCtx.DBConnection())
	bmUC := bmuc.NewBoardMemberUseCase(bmRepo, boardRepo, userRepo, appCtx.AsyncGroup())
	bmHdl := bmapi.NewBoardMemberHandler(bmUC)

	boardMemberRouter := boardRouter.Group("/:boardId/members")
	{
		boardMemberRouter.PATCH("", bmHdl.AddMember(appCtx))
		boardMemberRouter.PATCH("/:memberId", bmHdl.RemoveMember(appCtx))
		boardMemberRouter.GET("", bmHdl.GetBoardMembers(appCtx))
	}
}
