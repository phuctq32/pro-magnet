package authapi

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"net/http"
	"pro-magnet/common"
	"pro-magnet/components/appcontext"
	"strings"
	"time"
)

func (hdl *authHandler) LoginWithGoogle(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		data, err := hdl.uc.GoogleOAuthData()
		if err != nil {
			panic(err)
		}
		log.Debug().Str("state", data.State).Msg("")

		exp := time.Now().Add(time.Minute * 30)
		http.SetCookie(c.Writer, &http.Cookie{Name: "ggoauthstate", Value: data.State, Expires: exp})
		c.Redirect(http.StatusTemporaryRedirect, data.Url)
	}
}

func (hdl *authHandler) LoginWithGoogleCallback(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		oauthState, err := c.Request.Cookie("ggoauthstate")
		if err != nil {
			panic(common.NewBadRequestErr(errors.New("can not get oauth state")))
		}

		state := c.Query("state")
		log.Debug().Str("cookieState", oauthState.Value).Str("queryState", state).Msg("")
		if oauthState.Value != state {
			panic(common.NewBadRequestErr(errors.New("invalid google oauth state")))
		}

		code, ok := c.GetQuery("code")
		if !ok || strings.TrimSpace(code) == "" {
			panic(common.NewBadRequestErr(errors.New("can not get google oauth code")))
		}

		res, err := hdl.uc.LoginWithGoogle(c.Request.Context(), code)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewResponse("login successfully", res))
	}
}
