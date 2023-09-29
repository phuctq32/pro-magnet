package usermodel

import "time"

const UserCollectionName string = "users"

type User struct {
	Id          *string   `json:"_id,omitempty" bson:"_id,omitempty"`
	CreatedAt   time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt" bson:"updatedAt"`
	Email       string    `json:"email" bson:"email"`
	Name        string    `json:"name" bson:"name"`
	Password    string    `json:"password" bson:"password"`
	IsVerified  bool      `json:"isVerified" bson:"isVerified"`
	Avatar      string    `json:"avatar" bson:"avatar"`
	PhoneNumber string    `json:"phoneNumber" bson:"phoneNumber"`
	Birthday    time.Time `json:"birthday" bson:"birthday"`
}
