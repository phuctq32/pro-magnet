package cardapi

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
	"pro-magnet/common"
	"pro-magnet/components/appcontext"
	cardmodel "pro-magnet/modules/card/model"
	"reflect"
	"strings"
)

func (hdl *cardHandler) UpdateCardById(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		cardIdData := struct {
			CardId string `json:"cardId" validate:"required,mongodb"`
		}{
			CardId: strings.TrimSpace(c.Param("cardId")),
		}

		if errs := appCtx.Validator().Validate(&cardIdData); errs != nil {
			panic(common.NewValidationErrors(errs))
		}

		var data cardmodel.CardUpdate

		if err := c.ShouldBind(&data); err != nil {
			panic(common.NewBadRequestErr(err))
		}

		if errs := appCtx.Validator().Validate(&data); errs != nil {
			panic(common.NewValidationErrors(errs))
		}

		if reflect.ValueOf(data).IsZero() {
			panic(common.NewBadRequestErr(errors.New("invalid request")))
		}

		userId := c.MustGet(common.RequesterKey).(common.Requester).UserId()

		card, err := hdl.uc.UpdateCardById(c.Request.Context(), userId, cardIdData.CardId, &data)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusCreated, common.NewResponse("updated card successfully", card))
	}
}
