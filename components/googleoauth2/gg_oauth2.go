package ggoauth2

import (
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"io"
	"net/http"
	"net/url"
	ggoauthmodel "pro-magnet/components/googleoauth2/model"
	"time"
)

type googleOauth2 struct {
	cfg *oauth2.Config
}

func NewGoogleOAuth2(clientId, clientSecret, redirectUrl string) GoogleOAuth {
	return &googleOauth2{cfg: &oauth2.Config{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		Endpoint:     google.Endpoint,
		RedirectURL:  redirectUrl,
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
			"https://www.googleapis.com/auth/user.birthday.read",
			"https://www.googleapis.com/auth/user.phonenumbers.read",
		},
	}}
}

func (ggo *googleOauth2) AuthURL(state string) string {
	return ggo.cfg.AuthCodeURL(state)
}

func (ggo *googleOauth2) UserInfoFromCode(ctx context.Context, code string) (*ggoauthmodel.User, error) {
	token, _ := ggo.cfg.Exchange(ctx, code)

	req, _ := http.NewRequestWithContext(ctx, "GET", "https://people.googleapis.com/v1/people/me", nil)
	reqQuery := url.Values{}
	reqQuery.Add("personFields", "emailAddresses")
	reqQuery.Add("personFields", "names")
	reqQuery.Add("personFields", "photos")
	reqQuery.Add("personFields", "birthdays")
	reqQuery.Add("personFields", "phoneNumbers")
	req.URL.RawQuery = reqQuery.Encode()

	// Set access token for request
	token.SetAuthHeader(req)

	res, _ := (&http.Client{}).Do(req)
	defer res.Body.Close()

	var resBody ggoauthmodel.GoogleUserResponse
	b, _ := io.ReadAll(res.Body)
	if err := json.Unmarshal(b, &resBody); err != nil {
		return nil, errors.New("can not unmarshal response body")
	}

	// Convert response data to google user
	ggUser := &ggoauthmodel.User{}

	if resBody.EmailAddresses != nil {
		ggUser.Email = resBody.EmailAddresses[0].Value
	}

	if resBody.Names != nil {
		ggUser.Name = resBody.Names[0].DisplayName
	}

	if resBody.Photos != nil {
		ggUser.Avatar = resBody.Photos[0].URL
	}

	if resBody.Birthdays != nil {
		birthday := time.Date(
			resBody.Birthdays[0].Date.Year,
			time.Month(resBody.Birthdays[0].Date.Month),
			resBody.Birthdays[0].Date.Day,
			0, 0, 0, 0, time.Local,
		)

		ggUser.Birthday = &birthday
	}

	if resBody.PhoneNumbers != nil {
		*ggUser.Phonenumber = resBody.PhoneNumbers[0].Value
	}

	log.Debug().Interface("gguser", *ggUser).Msg("")

	return ggUser, nil
}
