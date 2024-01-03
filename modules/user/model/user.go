package usermodel

import "time"

const UserCollectionName string = "users"

type UserType string

const (
	InternalUser UserType = "INTERNAL"
	GoogleUser   UserType = "GOOGLE_USER"
	FacebookUser UserType = "FACEBOOK_USER"
)

type User struct {
	Id          *string    `json:"_id,omitempty" bson:"_id,omitempty"`
	CreatedAt   *time.Time `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt   *time.Time `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
	Email       *string    `json:"email,omitempty" bson:"email,omitempty"`
	Name        string     `json:"name" bson:"name"`
	Password    *string    `json:"-" bson:"password,omitempty"`
	IsVerified  *bool      `json:"isVerified,omitempty" bson:"isVerified,omitempty"`
	Avatar      string     `json:"avatar" bson:"avatar"`
	PhoneNumber *string    `json:"phoneNumber,omitempty" bson:"phoneNumber,omitempty"`
	Birthday    *time.Time `json:"birthday,omitempty" bson:"birthday,omitempty"`
	Type        *UserType  `json:"type,omitempty" bson:"type,omitempty"`
	Skills      []string   `json:"skills,omitempty" bson:"skills,omitempty"`
}

func (u *User) UserId() string {
	return *u.Id
}
