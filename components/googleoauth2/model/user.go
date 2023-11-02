package ggoauthmodel

import "time"

type User struct {
	Email       string
	Name        string
	Avatar      string
	Birthday    *time.Time
	Phonenumber *string
}
