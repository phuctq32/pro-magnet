package columnapi

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pro-magnet/common"
	"pro-magnet/components/appcontext"
	columnmodel "pro-magnet/modules/column/model"
)

func (hdl *columnHandler) CreateColumn(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data columnmodel.ColumnCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(common.NewBadRequestErr(err))
		}

		if errs := appCtx.Validator().Validate(&data); errs != nil {
			panic(common.NewValidationErrors(errs))
		}

		userId := c.MustGet(common.RequesterKey).(common.Requester).UserId()

		col, err := hdl.uc.CreateColumn(c.Request.Context(), userId, &data)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusCreated, common.NewResponse("successfully created column", col))
	}
}
