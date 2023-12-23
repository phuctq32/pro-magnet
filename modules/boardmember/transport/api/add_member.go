package bmapi

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pro-magnet/common"
	"pro-magnet/components/appcontext"
	bmmodel "pro-magnet/modules/boardmember/model"
	"strings"
)

func (hdl *boardMemberHandler) AddMember(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data *bmmodel.AddBoardMembers
		if err := c.ShouldBind(&data); err != nil {
			panic(common.NewBadRequestErr(err))
		}
		data.BoardId = strings.TrimSpace(c.Param("boardId"))

		if errs := appCtx.Validator().Validate(data); errs != nil {
			panic(common.NewValidationErrors(errs))
		}

		requesterId := c.MustGet(common.RequesterKey).(common.Requester).UserId()

		if err := hdl.uc.AddMember(c.Request.Context(), requesterId, data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewResponse("added user to board", nil))
	}
}
