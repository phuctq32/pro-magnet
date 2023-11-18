package cardchecklistapi

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pro-magnet/common"
	"pro-magnet/components/appcontext"
	cardchecklistmodel "pro-magnet/modules/cardchecklist/model"
	"strings"
)

func (hdl *cardChecklistHandler) CreateChecklist(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		cardIdData := struct {
			CardId string `json:"cardId" validate:"required,mongodb"`
		}{
			CardId: strings.TrimSpace(c.Param("cardId")),
		}

		if errs := appCtx.Validator().Validate(&cardIdData); errs != nil {
			panic(common.NewValidationErrors(errs))
		}

		var data cardchecklistmodel.CardChecklist

		if err := c.ShouldBind(&data); err != nil {
			panic(common.NewBadRequestErr(err))
		}

		if errs := appCtx.Validator().Validate(&data); errs != nil {
			panic(common.NewValidationErrors(errs))
		}

		if err := hdl.uc.CreateChecklist(c.Request.Context(), cardIdData.CardId, &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewResponse("successfully created checklist", nil))
	}
}
