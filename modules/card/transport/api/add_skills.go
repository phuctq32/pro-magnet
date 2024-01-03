package cardapi

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pro-magnet/common"
	"pro-magnet/components/appcontext"
	"strings"
)

func (hdl *cardHandler) AddSkills(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		requesterId := c.MustGet(common.RequesterKey).(common.Requester).UserId()

		data := struct {
			CardId string   `json:"cardId" validate:"required,mongodb"`
			Skills []string `json:"skills" validate:"required,min=1,dive,required"`
		}{
			CardId: strings.TrimSpace(c.Param("cardId")),
		}

		if err := c.ShouldBind(&data); err != nil {
			panic(common.NewBadRequestErr(err))
		}

		if errs := appCtx.Validator().Validate(&data); errs != nil {
			panic(common.NewValidationErrors(errs))
		}

		if err := hdl.uc.AddSkils(c.Request.Context(), requesterId, data.CardId, data.Skills); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewResponse("added card skills", nil))
	}
}
