package recomapi

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pro-magnet/common"
	"pro-magnet/components/appcontext"
	"strings"
)

func (hdl *recomHandler) GetRecommendedUsersForCard(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		requesterId := c.MustGet(common.RequesterKey).(common.Requester).UserId()

		data := struct {
			CardId string `json:"cardId" validate:"required,mongodb"`
		}{
			CardId: strings.TrimSpace(c.Param("cardId")),
		}

		if errs := appCtx.Validator().Validate(&data); errs != nil {
			panic(common.NewValidationErrors(errs))
		}

		users, err := hdl.uc.GetRecommendedUsersForCard(c.Request.Context(), requesterId, data.CardId, 5)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewResponse("", users))
	}
}
