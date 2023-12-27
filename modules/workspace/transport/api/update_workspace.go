package wsapi

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pro-magnet/common"
	"pro-magnet/components/appcontext"
	wsmodel "pro-magnet/modules/workspace/model"
	"strings"
)

func (hdl *wsHandler) UpdateWorkspace(appCtx appcontext.AppContext) gin.HandlerFunc {
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

		var updateData wsmodel.WorkspaceUpdate
		if err := c.ShouldBind(&updateData); err != nil {
			panic(common.NewBadRequestErr(err))
		}
		if errs := appCtx.Validator().Validate(&updateData); errs != nil {
			panic(common.NewValidationErrors(errs))
		}

		if err := hdl.uc.UpdateWorkspace(c.Request.Context(), requesterId, data.WorkspaceId, &updateData); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewResponse("updated workspace", nil))
	}
}
