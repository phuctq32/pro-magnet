package authapi

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pro-magnet/common"
	"pro-magnet/components/appcontext"
)

func (hdl *authHandler) Refresh(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		data := struct {
			RefreshToken string `json:"refreshToken"`
		}{}

		if err := c.ShouldBind(&data); err != nil {
			panic(common.NewBadRequestErr(err))
		}

		accessToken, err := hdl.uc.RefreshAccessToken(c.Request.Context(), data.RefreshToken)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewResponse("", map[string]string{
			"accessToken": *accessToken,
		}))
	}
}
