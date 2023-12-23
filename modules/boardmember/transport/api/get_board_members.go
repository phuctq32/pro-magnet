package bmapi

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pro-magnet/common"
	"pro-magnet/components/appcontext"
	"strings"
)

func (hdl *boardMemberHandler) GetBoardMembers(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		data := struct {
			BoardId string `json:"boardId" validate:"required,mongodb"`
		}{
			BoardId: strings.TrimSpace(c.Param("boardId")),
		}

		if errs := appCtx.Validator().Validate(&data); errs != nil {
			panic(common.NewValidationErrors(errs))
		}

		requesterId := c.MustGet(common.RequesterKey).(common.Requester).UserId()

		users, err := hdl.uc.GetBoardMembers(c.Request.Context(), requesterId, data.BoardId)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewResponse("", users))
	}
}
