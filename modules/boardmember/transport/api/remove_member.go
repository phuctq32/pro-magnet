package bmapi

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pro-magnet/common"
	"pro-magnet/components/appcontext"
	bmmodel "pro-magnet/modules/boardmember/model"
)

func (hdl *boardMemberHandler) RemoveMember(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data bmmodel.BoardMember
		data.BoardId = c.Param("boardId")
		data.UserId = c.Param("memberId")

		if errs := appCtx.Validator().Validate(&data); errs != nil {
			panic(common.NewValidationErrors(errs))
		}

		requesterId := c.MustGet(common.RequesterKey).(common.Requester).UserId()

		if err := hdl.uc.RemoveMember(c.Request.Context(), requesterId, &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewResponse("Removed member", nil))
	}
}
