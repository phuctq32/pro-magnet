package userapi

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pro-magnet/common"
	"pro-magnet/components/appcontext"
	"strings"
)

func (hdl *userHandler) RemoveSkill(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		requesterId := c.MustGet(common.RequesterKey).(common.Requester).UserId()

		data := struct {
			Skill string `json:"skill" validate:"required"`
		}{
			Skill: strings.TrimSpace(c.Query("value")),
		}

		if errs := appCtx.Validator().Validate(&data); errs != nil {
			panic(common.NewValidationErrors(errs))
		}

		if err := hdl.uc.RemoveSkill(c.Request.Context(), requesterId, data.Skill); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewResponse("removed skill", nil))
	}
}
