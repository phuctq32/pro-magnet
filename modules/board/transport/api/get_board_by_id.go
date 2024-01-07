package boardapi

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pro-magnet/common"
	"pro-magnet/components/appcontext"
	"strings"
)

func (hdl *boardHandler) GetBoardById(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		requesterId := c.MustGet(common.RequesterKey).(common.Requester).UserId()

		data := struct {
			BoardId string `json:"boardId" validate:"required,mongodb"`
		}{
			BoardId: strings.TrimSpace(c.Param("boardId")),
		}
		if errs := appCtx.Validator().Validate(&data); errs != nil {
			panic(common.NewValidationErrors(errs))
		}

		labelIdsData := struct {
			LabelIds []string `json:"labelIds" validate:"required,dive,mongodb"`
		}{
			LabelIds: c.QueryArray("labelIds"),
		}

		if len(labelIdsData.LabelIds) > 0 {
			if errs := appCtx.Validator().Validate(&labelIdsData); errs != nil {
				panic(common.NewValidationErrors(errs))
			}
		} else {
			labelIdsData.LabelIds = nil
		}

		board, err := hdl.uc.GetBoardById(c.Request.Context(), requesterId, data.BoardId, labelIdsData.LabelIds)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewResponse("", board))
	}
}
