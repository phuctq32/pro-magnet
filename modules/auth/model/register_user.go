package authmodel

import "time"

const DefaultAvatarUrl = "https://simulacionymedicina.es/wp-content/uploads/2015/11/default-avatar-300x300-1.jpg"

type RegisterUser struct {
	Email           string    `json:"email" validate:"required,email"`
	Name            string    `json:"name" validate:"required"`
	Password        string    `json:"password" validate:"required,min=6,alphanumunicode"`
	ConfirmPassword string    `json:"confirmPassword" validate:"required,eqfield=Password"`
	PhoneNumber     string    `json:"phoneNumber" validate:""`
	Birthday        time.Time `json:"birthday" validate:"required"`
}
