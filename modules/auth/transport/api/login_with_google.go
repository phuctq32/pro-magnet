package authapi

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pro-magnet/common"
	"pro-magnet/components/appcontext"
	"strings"
)

func (hdl *authHandler) LoginWithGoogle(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		//oauthState, err := c.Request.Cookie("ggoauthstate")
		//if err != nil {
		//	panic(common.NewBadRequestErr(errors.New("can not get oauth state")))
		//}
		//
		//state := c.Query("state")
		//log.Debug().Str("cookieState", oauthState.Value).Str("queryState", state).Msg("")
		//if oauthState.Value != state {
		//	panic(common.NewBadRequestErr(errors.New("invalid google oauth state")))
		//}

		data := struct {
			Code string `json:"code" validate:"required"`
		}{}

		if err := c.ShouldBind(&data); err != nil {
			panic(common.NewBadRequestErr(err))
		}

		if errs := appCtx.Validator().Validate(&data); errs != nil {
			panic(common.NewValidationErrors(errs))
		}

		res, err := hdl.uc.LoginWithGoogle(c.Request.Context(), strings.TrimSpace(data.Code))
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewResponse("login successfully", res))
	}
}
