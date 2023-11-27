package usermodel

type UserChangePassword struct {
	CurrentPassword string `json:"currentPassword" validate:"required"`
	Password        string `json:"newPassword" validate:"required,min=6,alphanumunicode,nefield=CurrentPassword"`
	ConfirmPassword string `json:"confirmNewPassword" validate:"required,eqfield=Password"`
}
