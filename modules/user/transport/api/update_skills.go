package userapi

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pro-magnet/common"
	"pro-magnet/components/appcontext"
)

func (hdl *userHandler) UpdateSkills(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		requesterId := c.MustGet(common.RequesterKey).(common.Requester).UserId()

		data := struct {
			Skills []string `json:"skills" validate:"required,min=1,dive,required"`
		}{}

		if err := c.ShouldBind(&data); err != nil {
			panic(common.NewBadRequestErr(err))
		}

		if errs := appCtx.Validator().Validate(&data); errs != nil {
			panic(common.NewValidationErrors(errs))
		}

		if err := hdl.uc.UpdateSkills(c.Request.Context(), requesterId, data.Skills); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewResponse("updated user skills", nil))
	}
}
