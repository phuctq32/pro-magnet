package routesv1

import (
	"github.com/gin-gonic/gin"
	"pro-magnet/components/appcontext"
	hasher2 "pro-magnet/components/hasher"
	"pro-magnet/middlewares"
	userrepo "pro-magnet/modules/user/repository/mongo"
	userapi "pro-magnet/modules/user/transport/api"
	useruc "pro-magnet/modules/user/usecase"
)

func NewUserRouter(appCtx appcontext.AppContext, router *gin.RouterGroup) {
	userRepo := userrepo.NewUserRepository(appCtx.DBConnection())
	hasher := hasher2.NewBcryptHash(10)
	userUC := useruc.NewUserUseCase(userRepo, hasher)
	userHdl := userapi.NewUserHandler(userUC)

	userRouter := router.Group("/users", middlewares.Authorize(appCtx))
	{
		userRouter.PATCH("/me", userHdl.UpdateUser(appCtx))
		userRouter.GET("/me", userHdl.GetProfile(appCtx))
		userRouter.PATCH("/me/change-password", userHdl.ChangePassword(appCtx))
	}
}
