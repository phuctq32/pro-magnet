package caapi

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pro-magnet/common"
	"pro-magnet/components/appcontext"
	camodel "pro-magnet/modules/cardattachment/model"
	"strings"
)

func (hdl *cardAttachmentHandler) AddCardAttachment(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		cardIdData := struct {
			CardId string `json:"cardId" validate:"required,mongodb"`
		}{
			CardId: strings.TrimSpace(c.Param("cardId")),
		}

		if errs := appCtx.Validator().Validate(&cardIdData); errs != nil {
			panic(common.NewValidationErrors(errs))
		}

		var data camodel.CardAttachment
		data.CardId = cardIdData.CardId

		if err := c.ShouldBind(&data); err != nil {
			panic(common.NewBadRequestErr(err))
		}

		if errs := appCtx.Validator().Validate(&data); errs != nil {
			panic(common.NewValidationErrors(errs))
		}

		ca, err := hdl.uc.AddCardAttachment(c.Request.Context(), &data)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewResponse("added card attachment successfuly", ca))
	}
}
