package boardapi

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
	"pro-magnet/common"
	"pro-magnet/components/appcontext"
	boardmodel "pro-magnet/modules/board/model"
	"reflect"
	"strings"
)

func (hdl *boardHandler) UpdateBoard(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		requesterId := c.MustGet(common.RequesterKey).(common.Requester).UserId()

		boardIdData := struct {
			BoardId string `json:"boardId" validate:"required,mongodb"`
		}{
			BoardId: strings.TrimSpace(c.Param("boardId")),
		}

		if errs := appCtx.Validator().Validate(&boardIdData); errs != nil {
			panic(common.NewValidationErrors(errs))
		}

		var data boardmodel.BoardUpdate

		if err := c.ShouldBind(&data); err != nil {
			panic(common.NewBadRequestErr(err))
		}

		if errs := appCtx.Validator().Validate(&data); errs != nil {
			panic(common.NewValidationErrors(errs))
		}
		if reflect.ValueOf(data).IsZero() {
			panic(common.NewBadRequestErr(errors.New("invalid request")))
		}

		if err := hdl.uc.UpdateBoard(
			c.Request.Context(), requesterId,
			boardIdData.BoardId, &data,
		); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewResponse("updated board", nil))
	}
}
