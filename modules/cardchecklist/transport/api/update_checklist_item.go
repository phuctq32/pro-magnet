package cardchecklistapi

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pro-magnet/common"
	"pro-magnet/components/appcontext"
	cardchecklistmodel "pro-magnet/modules/cardchecklist/model"
	"strings"
)

func (hdl *cardChecklistHandler) UpdateChecklistItem(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		cardIdData := struct {
			CardId      string `json:"cardId" validate:"required,mongodb"`
			ChecklistId string `json:"checklistId" validate:"required,mongodb"`
			ItemId      string `json:"itemId" validate:"required,mongodb"`
		}{
			CardId:      strings.TrimSpace(c.Param("cardId")),
			ChecklistId: strings.TrimSpace(c.Param("checklistId")),
			ItemId:      strings.TrimSpace(c.Param("itemId")),
		}
		if errs := appCtx.Validator().Validate(&cardIdData); errs != nil {
			panic(common.NewValidationErrors(errs))
		}

		var data cardchecklistmodel.ChecklistItemUpdate
		if err := c.ShouldBind(&data); err != nil {
			panic(common.NewBadRequestErr(err))
		}

		if errs := appCtx.Validator().Validate(&data); errs != nil {
			panic(common.NewValidationErrors(errs))
		}

		if err := hdl.uc.UpdateChecklistItem(
			c.Request.Context(),
			cardIdData.CardId,
			cardIdData.ChecklistId,
			cardIdData.ItemId,
			&data,
		); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewResponse("updated checklist item", nil))
	}
}
