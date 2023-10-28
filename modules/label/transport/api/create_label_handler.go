package labelapi

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pro-magnet/common"
	"pro-magnet/components/appcontext"
	labelmodel "pro-magnet/modules/label/model"
)

func (hdl *labelHandler) CreateLabel(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data labelmodel.LabelCreation

		if err := c.ShouldBind(&data); err != nil {
			panic(common.NewBadRequestErr(err))
		}

		if errs := appCtx.Validator().Validate(&data); errs != nil {
			panic(common.NewValidationErrors(errs))
		}

		label, err := hdl.uc.CreateLabel(c.Request.Context(), &data)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewResponse("successfully created label", label))
	}
}
