package common

const RequesterKey = "requester"

type Requester interface {
	UserId() string
}
