package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/pkg/errors"
	"time"
)

type claims struct {
	Payload
	jwt.RegisteredClaims
}

type jwtProvider struct{}

func NewJwtProvider() *jwtProvider {
	return &jwtProvider{}
}

func (jp *jwtProvider) Generate(payload *Payload, secretKey string, expiry int) (*string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims{
		*payload,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Local().Add(time.Minute * time.Duration(expiry))),
			IssuedAt:  jwt.NewNumericDate(time.Now().Local()),
		},
	})

	token, err := t.SignedString([]byte(secretKey))
	if err != nil {
		return nil, errors.New("sign token failed")
	}

	return &token, nil
}

func (jp *jwtProvider) Validate(tokenStr string, secretKey string) (*Payload, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	// In this parse function, error returned is nil only if token.Valid == false
	// If error occurred, token must be invalid
	if err != nil || !token.Valid {
		return nil, err
	}

	claims, ok := token.Claims.(*claims)
	if !ok {
		return nil, errors.New("can not assert token claims")
	}

	return &claims.Payload, nil
}
