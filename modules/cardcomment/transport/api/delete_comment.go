package cardcommentapi

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pro-magnet/common"
	"pro-magnet/components/appcontext"
	"strings"
)

func (hdl *cardCommentHandler) DeleteComment(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		requesterId := c.MustGet(common.RequesterKey).(common.Requester).UserId()

		idData := struct {
			CardId    string `json:"cardId" validate:"required,mongodb"`
			CommentId string `json:"checklistId" validate:"required,mongodb"`
		}{
			CardId:    strings.TrimSpace(c.Param("cardId")),
			CommentId: strings.TrimSpace(c.Param("commentId")),
		}
		if errs := appCtx.Validator().Validate(&idData); errs != nil {
			panic(common.NewValidationErrors(errs))
		}

		if err := hdl.uc.DeleteCardComment(
			c.Request.Context(), requesterId,
			idData.CardId, idData.CommentId,
		); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewResponse("removed comment", nil))
	}
}
