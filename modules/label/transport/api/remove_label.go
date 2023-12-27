package labelapi

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pro-magnet/common"
	"pro-magnet/components/appcontext"
	"strings"
)

func (hdl *labelHandler) RemoveLabel(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		requesterId := c.MustGet(common.RequesterKey).(common.Requester).UserId()

		labelIdData := struct {
			LabelId string `json:"labelId" validate:"required,mongodb"`
		}{
			LabelId: strings.TrimSpace(c.Param("labelId")),
		}
		if errs := appCtx.Validator().Validate(&labelIdData); errs != nil {
			panic(common.NewValidationErrors(errs))
		}

		cardIdData := struct {
			CardId string `json:"cardId" validate:"required,mongodb"`
		}{
			CardId: strings.TrimSpace(c.Query("cardId")),
		}
		if strings.TrimSpace(cardIdData.CardId) != "" {
			if errs := appCtx.Validator().Validate(&labelIdData); errs != nil {
				panic(common.NewValidationErrors(errs))
			}

			if err := hdl.uc.RemoveLabelFromCard(
				c.Request.Context(), requesterId,
				cardIdData.CardId, labelIdData.LabelId,
			); err != nil {
				panic(err)
			}
			c.JSON(http.StatusOK, common.NewResponse("removed label from card", nil))
		} else {
			if err := hdl.uc.RemoveLabelFromBoard(
				c.Request.Context(),
				requesterId, labelIdData.LabelId,
			); err != nil {
				panic(err)
			}

			c.JSON(http.StatusOK, common.NewResponse("removed label from board", nil))
		}
	}
}
