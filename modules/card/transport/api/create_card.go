package cardapi

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pro-magnet/common"
	"pro-magnet/components/appcontext"
	cardmodel "pro-magnet/modules/card/model"
)

func (hdl *cardHandler) CreateCard(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data cardmodel.CardCreation

		if err := c.ShouldBind(&data); err != nil {
			panic(common.NewBadRequestErr(err))
		}

		if errs := appCtx.Validator().Validate(&data); errs != nil {
			panic(common.NewValidationErrors(errs))
		}

		userId := c.MustGet(common.RequesterKey).(common.Requester).UserId()

		card, err := hdl.uc.CreateCard(c.Request.Context(), userId, &data)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusCreated, common.NewResponse("created card successfully", card))
	}
}
