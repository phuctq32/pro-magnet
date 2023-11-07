package authapi

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
	"pro-magnet/common"
	"pro-magnet/components/appcontext"
	authmodel "pro-magnet/modules/auth/model"
	"strings"
)

func (hdl *authHandler) ResetPassword(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data authmodel.ResetPasswordUser

		resetToken, ok := c.GetQuery("resetToken")
		if !ok {
			panic(common.NewBadRequestErr(errors.New("can not get reset token")))
		}
		resetToken = strings.TrimSpace(resetToken)

		if err := c.ShouldBind(&data); err != nil {
			panic(common.NewBadRequestErr(err))
		}

		if errs := appCtx.Validator().Validate(&data); errs != nil {
			panic(common.NewValidationErrors(errs))
		}

		if err := hdl.uc.ResetPassword(c.Request.Context(), resetToken, data.Password); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewResponse("reset password successfully", nil))
	}
}
