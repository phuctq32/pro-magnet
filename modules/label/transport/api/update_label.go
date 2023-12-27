package labelapi

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pro-magnet/common"
	"pro-magnet/components/appcontext"
	labelmodel "pro-magnet/modules/label/model"
	"strings"
)

func (hdl *labelHandler) UpdateLabel(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		requesterId := c.MustGet(common.RequesterKey).(common.Requester).UserId()

		data := struct {
			LabelId string `json:"labelId" validate:"required,mongodb"`
		}{
			LabelId: strings.TrimSpace(c.Param("labelId")),
		}

		if errs := appCtx.Validator().Validate(&data); errs != nil {
			panic(common.NewValidationErrors(errs))
		}

		var updateData labelmodel.LabelUpdate
		if err := c.ShouldBind(&updateData); err != nil {
			panic(common.NewBadRequestErr(err))
		}
		if errs := appCtx.Validator().Validate(&updateData); errs != nil {
			panic(common.NewValidationErrors(errs))
		}

		if err := hdl.uc.UpdateLabel(c.Request.Context(), requesterId, data.LabelId, &updateData); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewResponse("updated label", nil))
	}
}
