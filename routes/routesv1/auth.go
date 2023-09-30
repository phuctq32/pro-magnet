package routesv1

import (
	"github.com/gin-gonic/gin"
	"pro-magnet/components/appcontext"
	hasher2 "pro-magnet/components/hasher"
	"pro-magnet/components/jwt"
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
	jwtProvider := jwt.NewJwtProvider()

	authUC := authuc.NewAuthUseCase(
		userRepo,
		authRedisRepo,
		hasher,
		sgMailer,
		jwtProvider,
		configs.EnvConfigs.AccessSecret(),
		configs.EnvConfigs.RefreshSecret(),
		configs.EnvConfigs.AccessTokenExpiry(),
		configs.EnvConfigs.RefreshTokenExpiry(),
	)

	authHdl := authapi.NewAuthHandler(authUC)

	// setup routes
	authRouter := router.Group("/auth")
	{
		authRouter.POST("/register", authHdl.Register(appCtx))
		authRouter.POST("/verify", authHdl.Verify(appCtx))
		authRouter.POST("/send-verification-email", authHdl.SendVerificationEmail(appCtx))
		authRouter.POST("/login", authHdl.Login(appCtx))
	}
}
