package authapi

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pro-magnet/common"
	"pro-magnet/components/appcontext"
	authmodel "pro-magnet/modules/auth/model"
)

func (hdl *authHandler) Login(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data authmodel.LoginUser
		if err := c.ShouldBind(&data); err != nil {
			panic(common.NewBadRequestErr(err))
		}

		if errs := appCtx.Validator().Validate(&data); errs != nil {
			panic(common.NewValidationErrors(errs))
		}

		tokenPair, err := hdl.uc.Login(c.Request.Context(), &data)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewResponse("", tokenPair))
	}
}
