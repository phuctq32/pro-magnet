package wsmemberapi

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pro-magnet/common"
	"pro-magnet/components/appcontext"
	"strings"
)

func (hdl *wsMemberHandler) RemoveMemberFromWorkspace(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		data := struct {
			WorkspaceId string `json:"workspaceId" validate:"required,mongodb"`
			MemberId    string `json:"memberId" validate:"required,mongodb"`
		}{
			WorkspaceId: strings.TrimSpace(c.Param("workspaceId")),
			MemberId:    strings.TrimSpace(c.Param("memberId")),
		}

		if errs := appCtx.Validator().Validate(&data); errs != nil {
			panic(common.NewValidationErrors(errs))
		}

		requesterId := c.MustGet(common.RequesterKey).(common.Requester).UserId()

		if err := hdl.uc.RemoveMember(
			c.Request.Context(), requesterId,
			data.WorkspaceId, data.MemberId,
		); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewResponse("removed user from workspace", nil))
	}
}
