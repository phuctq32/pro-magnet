package cardcommentapi

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pro-magnet/common"
	"pro-magnet/components/appcontext"
	cardcommentmodel "pro-magnet/modules/cardcomment/model"
	"strings"
)

func (hdl *cardCommentHandler) CreateCardComment(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		requesterId := c.MustGet(common.RequesterKey).(common.Requester).UserId()

		cardIdData := struct {
			CardId string `json:"cardId" validate:"required,mongodb"`
		}{
			CardId: strings.TrimSpace(c.Param("cardId")),
		}

		if errs := appCtx.Validator().Validate(&cardIdData); errs != nil {
			panic(common.NewValidationErrors(errs))
		}

		var data cardcommentmodel.CardCommentCreate
		if err := c.ShouldBind(&data); err != nil {
			panic(common.NewBadRequestErr(err))
		}

		if errs := appCtx.Validator().Validate(&data); errs != nil {
			panic(common.NewValidationErrors(errs))
		}

		data.UserId = requesterId
		if err := hdl.uc.CreateCardComment(c.Request.Context(), cardIdData.CardId, &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewResponse("created card comment", nil))
	}
}
