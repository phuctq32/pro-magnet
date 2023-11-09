package boardapi

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pro-magnet/common"
	"pro-magnet/components/appcontext"
	boardmodel "pro-magnet/modules/board/model"
)

func (hdl *boardHandler) CreateBoard(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data boardmodel.BoardCreation
		if err := c.ShouldBind(&data); err != nil {
			panic(common.NewBadRequestErr(err))
		}

		if errs := appCtx.Validator().Validate(&data); errs != nil {
			panic(common.NewValidationErrors(errs))
		}

		data.UserId = c.MustGet(common.RequesterKey).(common.Requester).UserId()

		board, err := hdl.uc.CreateBoard(c.Request.Context(), &data)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewResponse("created board successfully", board))
	}
}
