package tokens

import (
	"errors"
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func VerifyToken(token string) (*jwt.RegisteredClaims, error) {

	claims := &jwt.RegisteredClaims{}

	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, errors.New("invalid signature of jwt: " + err.Error())
		}
		return nil, errors.New("invalid jwt: " + err.Error())
	}

	if !tkn.Valid {
		fmt.Println("invalid jwt: ", tkn.Raw)
		return nil, errors.New("invalid jwt")
	}

	return claims, nil
}
