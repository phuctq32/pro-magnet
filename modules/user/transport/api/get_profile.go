package userapi

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pro-magnet/common"
	"pro-magnet/components/appcontext"
)

func (hdl *userHandler) GetProfile(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.MustGet(common.RequesterKey).(common.Requester).UserId()

		user, err := hdl.uc.GetUser(c.Request.Context(), userId)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewResponse("", user))
	}
}
