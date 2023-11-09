package authapi

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
	"pro-magnet/common"
	"pro-magnet/components/appcontext"
)

func (hdl *authHandler) Verify(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		verifiedToken, ok := c.GetQuery("token")
		if !ok {
			panic(common.NewBadRequestErr(errors.New("can not get verified verifiedToken")))
		}

		if err := hdl.uc.Verify(c.Request.Context(), verifiedToken); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewResponse("email verified", nil))
	}
}
