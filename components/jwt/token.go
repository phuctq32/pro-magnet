package jwt

type TokenProvider interface {
	Generate(payload *Payload, secretKey string, expiry int) (*string, error)
	Validate(token string, secretKey string) (*Payload, error)
}

type Payload struct {
	UserId string
}
