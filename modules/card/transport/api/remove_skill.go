package cardapi

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pro-magnet/common"
	"pro-magnet/components/appcontext"
	"strings"
)

func (hdl *cardHandler) RemoveSkill(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		requesterId := c.MustGet(common.RequesterKey).(common.Requester).UserId()

		data := struct {
			CardId string `json:"cardId" validate:"required,mongodb"`
			Skill  string `json:"skill" validate:"required"`
		}{
			Skill:  strings.TrimSpace(c.Query("value")),
			CardId: strings.TrimSpace(c.Param("cardId")),
		}

		if errs := appCtx.Validator().Validate(&data); errs != nil {
			panic(common.NewValidationErrors(errs))
		}

		if err := hdl.uc.RemoveSkill(c.Request.Context(), requesterId, data.CardId, data.Skill); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewResponse("removed skill", nil))
	}
}
