package v1

import (
	"github.com/gin-gonic/gin"
	"pro-magnet/components/appcontext"
	hasher2 "pro-magnet/components/hasher"
	"pro-magnet/components/mailer/sendgrid"
	"pro-magnet/configs"
	authrepo "pro-magnet/modules/auth/repository/redis"
	authapi "pro-magnet/modules/auth/transport/api"
	authuc "pro-magnet/modules/auth/usecase"
	userrepo "pro-magnet/modules/user/repository/mongo"
)

func NewAuthRouter(appCtx appcontext.AppContext, router *gin.RouterGroup) {
	// setup dependencies
	userRepo := userrepo.NewUserRepository(appCtx.DBConnection())
	authRedisRepo := authrepo.NewAuthRedisRepository(appCtx.RedisClient())
	hasher := hasher2.NewBcryptHash(10)
	sgMailer := sendgrid.NewSendGridProvider(configs.EnvConfigs.SendgridApiKey())
	authUC := authuc.NewAuthUseCase(userRepo, authRedisRepo, hasher, sgMailer)
	authHdl := authapi.NewAuthHandler(authUC)

	// setup routes
	router.POST("/register", authHdl.Register(appCtx))
	router.POST("/verify", authHdl.Verify(appCtx))
}
