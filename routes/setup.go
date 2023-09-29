package routes

import (
	"github.com/gin-gonic/gin"
	"pro-magnet/components/appcontext"
	v1 "pro-magnet/routes/v1"
)

func Setup(appCtx appcontext.AppContext, engine *gin.Engine) {
	v1.NewAuthRouter(appCtx, engine.Group("/auth"))
}
