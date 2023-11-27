package usermodel

import "time"

type UserUpdate struct {
	Name        *string    `json:"name,omitempty" bson:"name,omitempty" validate:"omitempty,required"`
	PhoneNumber *string    `json:"phoneNumber,omitempty" bson:"phoneNumber,omitempty" validate:"omitempty,required"`
	Birthday    *time.Time `json:"birthday,omitempty" bson:"birthday,omitempty" validate:"omitempty,required"`
}
