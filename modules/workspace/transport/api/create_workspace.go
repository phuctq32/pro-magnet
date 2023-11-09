package wsapi

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pro-magnet/common"
	"pro-magnet/components/appcontext"
	wsmodel "pro-magnet/modules/workspace/model"
)

func (hdl *wsHandler) CreateWorkspace(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data wsmodel.WorkspaceCreation

		if err := c.ShouldBind(&data); err != nil {
			panic(common.NewBadRequestErr(err, "can not get data"))
		}

		if errs := appCtx.Validator().Validate(&data); errs != nil {
			panic(common.NewValidationErrors(errs))
		}

		userId := c.MustGet(common.RequesterKey).(common.Requester).UserId()

		newWs, err := hdl.uc.CreateWorkspace(c.Request.Context(), userId, &data)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewResponse("created workspace successfully", newWs))
	}
}
