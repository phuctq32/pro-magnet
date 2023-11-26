package columnapi

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pro-magnet/common"
	"pro-magnet/components/appcontext"
	columnmodel "pro-magnet/modules/column/model"
	"strings"
)

func (hdl *columnHandler) UpdateColumn(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		ColumIdData := struct {
			ColumnId string `json:"columnId" validate:"required,mongodb"`
		}{
			ColumnId: strings.TrimSpace(c.Param("id")),
		}

		if errs := appCtx.Validator().Validate(&ColumIdData); errs != nil {
			panic(common.NewValidationErrors(errs))
		}

		var updateData columnmodel.ColumnUpdate
		if err := c.ShouldBind(&updateData); err != nil {
			panic(err)
		}

		if errs := appCtx.Validator().Validate(&updateData); errs != nil {
			panic(common.NewValidationErrors(errs))
		}

		if errs := appCtx.Validator().Validate(&ColumIdData); errs != nil {
			panic(common.NewValidationErrors(errs))
		}

		userId := c.MustGet(common.RequesterKey).(common.Requester).UserId()

		updatedCol, err := hdl.uc.UpdateColumn(
			c.Request.Context(), userId,
			ColumIdData.ColumnId, &updateData)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewResponse("removed column", updatedCol))
	}
}
