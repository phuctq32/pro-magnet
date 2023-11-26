package columnapi

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pro-magnet/common"
	"pro-magnet/components/appcontext"
	"strings"
)

func (hdl *columnHandler) RemoveColumn(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		data := struct {
			ColumnId string `json:"columnId" validate:"required,mongodb"`
		}{
			ColumnId: strings.TrimSpace(c.Param("id")),
		}

		if errs := appCtx.Validator().Validate(&data); errs != nil {
			panic(common.NewValidationErrors(errs))
		}

		userId := c.MustGet(common.RequesterKey).(common.Requester).UserId()

		if err := hdl.uc.RemoveColumn(c.Request.Context(), userId, data.ColumnId); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewResponse("removed column", nil))
	}
}
