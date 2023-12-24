package cardcommentapi

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pro-magnet/common"
	"pro-magnet/components/appcontext"
	cardcommentmodel "pro-magnet/modules/cardcomment/model"
	"strings"
)

func (hdl *cardCommentHandler) UpdateCardComment(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		requesterId := c.MustGet(common.RequesterKey).(common.Requester).UserId()

		idData := struct {
			CardId    string `json:"cardId" validate:"required,mongodb"`
			CommentId string `json:"commentId" validate:"required,mongodb"`
		}{
			CardId:    strings.TrimSpace(c.Param("cardId")),
			CommentId: strings.TrimSpace(c.Param("commentId")),
		}

		if errs := appCtx.Validator().Validate(&idData); errs != nil {
			panic(common.NewValidationErrors(errs))
		}

		var data cardcommentmodel.CardCommentUpdate
		if err := c.ShouldBind(&data); err != nil {
			panic(common.NewBadRequestErr(err))
		}

		if errs := appCtx.Validator().Validate(&data); errs != nil {
			panic(common.NewValidationErrors(errs))
		}

		if err := hdl.uc.UpdateCardComment(
			c.Request.Context(),
			requesterId, idData.CardId,
			idData.CommentId, &data,
		); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewResponse("updated card comment", nil))
	}
}
