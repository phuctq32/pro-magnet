package wsmemberapi

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pro-magnet/common"
	"pro-magnet/components/appcontext"
	wsmembermodel "pro-magnet/modules/workspacemember/model"
	"strings"
)

func (hdl *wsMemberHandler) AddWorkspaceMembers(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data wsmembermodel.WorkspaceMembersCreate
		if err := c.ShouldBind(&data); err != nil {
			panic(common.NewBadRequestErr(err))
		}
		data.WorkspaceId = strings.TrimSpace(c.Param("workspaceId"))

		if errs := appCtx.Validator().Validate(&data); errs != nil {
			panic(common.NewValidationErrors(errs))
		}

		requesterId := c.MustGet(common.RequesterKey).(common.Requester).UserId()

		if err := hdl.uc.AddMembers(c.Request.Context(), requesterId, &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewResponse("added users to workspace", nil))
	}
}
