package userapi

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pro-magnet/common"
	"pro-magnet/components/appcontext"
	"strings"
)

func (hdl *userHandler) GetUsersByEmailToAddToWorkspace(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		requesterId := c.MustGet(common.RequesterKey).(common.Requester).UserId()

		data := struct {
			WorkspaceId      string `json:"workspaceId" validate:"required,mongodb"`
			EmailSearchQuery string `json:"emailQ" validate:"required"`
		}{
			WorkspaceId:      strings.TrimSpace(c.Query("workspaceId")),
			EmailSearchQuery: strings.TrimSpace(c.Query("emailQ")),
		}

		if errs := appCtx.Validator().Validate(&data); errs != nil {
			panic(common.NewValidationErrors(errs))
		}

		users, err := hdl.uc.GetUsersToAddToWorkspace(
			c.Request.Context(), requesterId,
			data.WorkspaceId, data.EmailSearchQuery,
		)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewResponse("", users))
	}
}
