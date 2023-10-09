package routes

import (
	"github.com/gin-gonic/gin"
	"pro-magnet/components/appcontext"
	"pro-magnet/routes/routesv1"
)

func Setup(appCtx appcontext.AppContext, engine *gin.Engine) {
	v1 := engine.Group("api/v1")
	routesv1.NewAuthRouter(appCtx, v1)
	routesv1.NewWorkspaceRouter(appCtx, v1)
}
