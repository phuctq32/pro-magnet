package cardapi

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pro-magnet/common"
	"pro-magnet/components/appcontext"
	"strings"
)

func (hdl *cardHandler) RemoveCard(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		cardIdData := struct {
			CardId string `json:"cardId" validate:"required,mongodb"`
		}{
			CardId: strings.TrimSpace(c.Param("cardId")),
		}

		if errs := appCtx.Validator().Validate(&cardIdData); errs != nil {
			panic(common.NewValidationErrors(errs))
		}

		requesterId := c.MustGet(common.RequesterKey).(common.Requester).UserId()

		err := hdl.uc.RemoveCard(c.Request.Context(), requesterId, cardIdData.CardId)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusCreated, common.NewResponse("removed card", nil))
	}
}
