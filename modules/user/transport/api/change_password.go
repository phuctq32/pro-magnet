package userapi

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pro-magnet/common"
	"pro-magnet/components/appcontext"
	usermodel "pro-magnet/modules/user/model"
)

func (hdl *userHandler) ChangePassword(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.MustGet(common.RequesterKey).(common.Requester).UserId()

		var data usermodel.UserChangePassword

		if err := c.ShouldBind(&data); err != nil {
			panic(common.NewBadRequestErr(err))
		}

		if errs := appCtx.Validator().Validate(&data); errs != nil {
			panic(common.NewValidationErrors(errs))
		}

		if err := hdl.uc.ChangePassword(c.Request.Context(), userId, &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewResponse("changed password successfully", nil))
	}
}
