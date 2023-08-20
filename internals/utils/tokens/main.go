package tokens

import "github.com/golang-jwt/jwt/v5"

type TokenClaims struct {
	RegisteredClaims *jwt.RegisteredClaims
}
