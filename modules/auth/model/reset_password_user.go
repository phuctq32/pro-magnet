package authmodel

type ResetPasswordUser struct {
	Password        string `json:"password" validate:"required,min=6,alphanumunicode"`
	ConfirmPassword string `json:"confirmPassword" validate:"eqfield=Password"`
}
