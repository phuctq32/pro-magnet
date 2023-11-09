package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"pro-magnet/common"
)

func Recover() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.Header("Content-Type", "application/json")

				appErr, ok := err.(*common.AppError)
				if !ok {
					appErr = common.NewServerErr(err.(error))
				}

				log.Error().Err(appErr).Msg(appErr.Log)
				c.AbortWithStatusJSON(appErr.StatusCode, appErr)

				if gin.Mode() == gin.DebugMode {
					panic(appErr)
				}

				return
			}
		}()

		c.Next()
	}
}
