package searchapi

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pro-magnet/common"
	"pro-magnet/components/appcontext"
	"strings"
)

func (hdl *searchHandler) Search(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		requesterId := c.MustGet(common.RequesterKey).(common.Requester).UserId()

		data := struct {
			SearchTerm string `json:"searchTerm" validate:"required"`
		}{
			SearchTerm: strings.TrimSpace(c.Query("q")),
		}

		if errs := appCtx.Validator().Validate(&data); errs != nil {
			panic(common.NewValidationErrors(errs))
		}

		searchData, err := hdl.uc.Search(c.Request.Context(), requesterId, data.SearchTerm)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewResponse("", searchData))
	}
}
