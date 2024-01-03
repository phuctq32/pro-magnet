package wsapi

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pro-magnet/common"
	"pro-magnet/components/appcontext"
)

func (hdl *wsHandler) GetCurrentUserWorkspaces(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		requesterId := c.MustGet(common.RequesterKey).(common.Requester).UserId()

		workspaces, err := hdl.uc.GetCurrentUserWorkspaces(c.Request.Context(), requesterId)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewResponse("", workspaces))
	}
}
