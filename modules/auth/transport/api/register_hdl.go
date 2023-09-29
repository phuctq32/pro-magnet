package authapi

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pro-magnet/common"
	"pro-magnet/components/appcontext"
	authmodel "pro-magnet/modules/auth/model"
)

func (hdl *authHandler) Register(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data authmodel.RegisterUser

		if err := c.ShouldBind(&data); err != nil {
			panic(common.NewBadRequestErr(err))
		}

		// validate
		if errs := appCtx.Validator().Validate(&data); errs != nil {
			panic(common.NewValidationErrors(errs))
		}

		if err := hdl.uc.Register(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusCreated, common.NewResponse("created user successfully", nil))
	}
}
