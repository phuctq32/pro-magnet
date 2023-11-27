package routesv1

import (
	"github.com/gin-gonic/gin"
	"pro-magnet/components/appcontext"
	"pro-magnet/middlewares"
	userrepo "pro-magnet/modules/user/repository/mongo"
	userapi "pro-magnet/modules/user/transport/api"
	useruc "pro-magnet/modules/user/usecase"
)

func NewUserRouter(appCtx appcontext.AppContext, router *gin.RouterGroup) {
	userRepo := userrepo.NewUserRepository(appCtx.DBConnection())
	userUC := useruc.NewUserUseCase(userRepo)
	userHdl := userapi.NewUserHandler(userUC)

	userRouter := router.Group("/users", middlewares.Authorize(appCtx))
	{
		userRouter.GET("/me", userHdl.GetProfile(appCtx))
	}
}
