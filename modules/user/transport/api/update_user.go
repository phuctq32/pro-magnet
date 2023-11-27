package userapi

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pro-magnet/common"
	"pro-magnet/components/appcontext"
	usermodel "pro-magnet/modules/user/model"
)

func (hdl *userHandler) UpdateUser(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.MustGet(common.RequesterKey).(common.Requester).UserId()

		var data usermodel.UserUpdate

		if err := c.ShouldBind(&data); err != nil {
			panic(common.NewBadRequestErr(err))
		}

		if errs := appCtx.Validator().Validate(&data); errs != nil {
			panic(common.NewValidationErrors(errs))
		}

		updatedUser, err := hdl.uc.UpdateUser(c.Request.Context(), userId, &data)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewResponse("updated user successfully", updatedUser))
	}
}
