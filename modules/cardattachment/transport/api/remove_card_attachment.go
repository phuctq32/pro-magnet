package caapi

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pro-magnet/common"
	"pro-magnet/components/appcontext"
	"strings"
)

func (hdl *cardAttachmentHandler) RemoveCardAttachment(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		data := struct {
			CardId       string `json:"cardId" validate:"required,mongodb"`
			AttachmentId string `json:"attachmentId" validate:"required,mongodb"`
		}{
			CardId:       strings.TrimSpace(c.Param("cardId")),
			AttachmentId: strings.TrimSpace(c.Param("attachmentId")),
		}

		if errs := appCtx.Validator().Validate(&data); errs != nil {
			panic(common.NewValidationErrors(errs))
		}

		userId := c.MustGet(common.RequesterKey).(common.Requester).UserId()

		if err := hdl.uc.RemoveCardAttachment(c.Request.Context(), userId, data.CardId, data.AttachmentId); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewResponse("Deleted attachment", nil))
	}
}
