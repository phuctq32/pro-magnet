package cardapi

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pro-magnet/common"
	"pro-magnet/components/appcontext"
	"strings"
)

func (hdl *cardHandler) RemoveMemberFromCard(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		requesterId := c.MustGet(common.RequesterKey).(common.Requester).UserId()

		idData := struct {
			CardId   string `json:"cardId" validate:"required,mongodb"`
			MemberId string `json:"userId" validate:"required,mongodb"`
		}{
			CardId:   strings.TrimSpace(c.Param("cardId")),
			MemberId: strings.TrimSpace(c.Param("memberId")),
		}

		if errs := appCtx.Validator().Validate(&idData); errs != nil {
			panic(common.NewValidationErrors(errs))
		}

		if err := hdl.uc.RemoveMemberFromCard(
			c.Request.Context(),
			requesterId,
			idData.CardId,
			idData.MemberId,
		); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewResponse("removed user from card", nil))
	}
}
