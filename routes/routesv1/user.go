package routesv1

import (
	"github.com/gin-gonic/gin"
	"pro-magnet/components/appcontext"
	hasher2 "pro-magnet/components/hasher"
	"pro-magnet/middlewares"
	boardmemberrepo "pro-magnet/modules/boardmember/repository/mongo"
	"pro-magnet/modules/card/repository/mongo"
	userrepo "pro-magnet/modules/user/repository/mongo"
	userapi "pro-magnet/modules/user/transport/api"
	useruc "pro-magnet/modules/user/usecase"
	wsrepo "pro-magnet/modules/workspace/repository/mongo"
	wsmemberrepo "pro-magnet/modules/workspacemember/repository/mongo"
)

func NewUserRouter(appCtx appcontext.AppContext, router *gin.RouterGroup) {
	userRepo := userrepo.NewUserRepository(appCtx.DBConnection())
	cardRepo := mongo.NewCardRepository(appCtx.DBConnection())
	bmRepo := boardmemberrepo.NewBoardMemberRepository(appCtx.DBConnection())
	wsRepo := wsrepo.NewWorkspaceRepository(appCtx.DBConnection())
	wsMemberRepo := wsmemberrepo.NewWorkspaceMemberRepository(appCtx.DBConnection())
	hasher := hasher2.NewBcryptHash(10)
	userUC := useruc.NewUserUseCase(userRepo, cardRepo, bmRepo, wsRepo, wsMemberRepo, hasher)
	userHdl := userapi.NewUserHandler(userUC)

	userRouter := router.Group("/users", middlewares.Authorize(appCtx))
	{
		userRouter.PATCH("/me", userHdl.UpdateUser(appCtx))
		userRouter.GET("/me", userHdl.GetProfile(appCtx))
		userRouter.PATCH("/me/change-password", userHdl.ChangePassword(appCtx))

		userRouter.GET("/to-add-to-card", userHdl.GetUsersToAddToCard(appCtx))
		userRouter.GET("/to-add-to-workspace", userHdl.GetUsersByEmailToAddToWorkspace(appCtx))
	}
}
