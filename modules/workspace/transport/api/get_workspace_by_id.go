package wsapi

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pro-magnet/common"
	"pro-magnet/components/appcontext"
	"strings"
)

func (hdl *wsHandler) GetWorkspaceById(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		requesterId := c.MustGet(common.RequesterKey).(common.Requester).UserId()

		data := struct {
			WorkspaceId string `json:"workspaceId" validate:"required,mongodb"`
		}{
			WorkspaceId: strings.TrimSpace(c.Param("workspaceId")),
		}

		if errs := appCtx.Validator().Validate(&data); errs != nil {
			panic(common.NewValidationErrors(errs))
		}
		ws, err := hdl.uc.GetWorkspaceById(c.Request.Context(), requesterId, data.WorkspaceId)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewResponse("", ws))
	}
}
