package cardchecklistapi

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pro-magnet/common"
	"pro-magnet/components/appcontext"
	"strings"
)

func (hdl *cardChecklistHandler) DeleteChecklistItem(appCtx appcontext.AppContext) gin.HandlerFunc {
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

		if err := hdl.uc.DeleteChecklistItem(
			c.Request.Context(),
			cardIdData.CardId,
			cardIdData.ChecklistId,
			cardIdData.ItemId,
		); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewResponse("removed checklist item", nil))
	}
}
