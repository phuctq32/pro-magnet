package authapi

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pro-magnet/common"
	"pro-magnet/components/appcontext"
)

func (hdl *authHandler) ForgotPassword(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		data := struct {
			Email string `json:"email" validate:"required,email"`
		}{}

		if err := c.ShouldBind(&data); err != nil {
			panic(common.NewBadRequestErr(err))
		}

		if errs := appCtx.Validator().Validate(&data); errs != nil {
			panic(common.NewValidationErrors(errs))
		}

		if err := hdl.uc.ForgotPassword(c.Request.Context(), data.Email); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewResponse("The reset link was sent to your email", nil))
	}
}
