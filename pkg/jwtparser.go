package pkg

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	jwt.StandardClaims
	Username string `json:"username"`
}

func ValidateToken(t string, secret string) (claims jwt.MapClaims, err error) {
	// Parse jwt token
	token, err := jwt.ParseWithClaims(
		t,
		&claims,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		},
	)
	if err != nil || !token.Valid {
		return claims, fmt.Errorf("unable to parse token or it's invalid")
	}

	return claims, nil
}
