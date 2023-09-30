package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"pro-magnet/common"
	"pro-magnet/components/appcontext"
	"pro-magnet/components/jwt"
	"pro-magnet/configs"
	userrepo "pro-magnet/modules/user/repository/mongo"
	"strings"
)

func Authorize(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract token from header
		parts := strings.Split(c.GetHeader("Authorization"), " ")
		if parts[0] != "Bearer" || len(parts) < 2 || strings.TrimSpace(parts[1]) == "" {
			panic(common.NewUnauthorizedErr(errors.New("invalid header")))
		}
		tokenProvider := jwt.NewJwtProvider()

		payload, err := tokenProvider.Validate(parts[1], configs.EnvConfigs.AccessSecret())
		if err != nil {
			panic(common.NewUnauthorizedErr(err, "invalid token"))
		}

		userRepo := userrepo.NewUserRepository(appCtx.DBConnection())

		user, err := userRepo.FindById(c.Request.Context(), payload.UserId)
		if err != nil {
			panic(err)
		}

		c.Set(common.RequesterKey, user)
		c.Next()
	}
}
