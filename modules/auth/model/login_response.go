package authmodel

import usermodel "pro-magnet/modules/user/model"

type LoginResponse struct {
	AccessToken  string          `json:"accessToken"`
	RefreshToken string          `json:"refreshToken"`
	User         *usermodel.User `json:"user"`
}
