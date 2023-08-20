package tokens

import (
	"auth-service/internals/utils/exceptions"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateRefreshToken(userId int64) (string, *exceptions.AuthError) {

	key := []byte(os.Getenv("JWT_SECRET"))

	exp := time.Now().UTC().Add(24 * time.Hour * 30) // 30 days

	claims := jwt.RegisteredClaims{
		Issuer:    "kddw-auth-service",
		Subject:   strconv.FormatInt(userId, 10),
		ExpiresAt: jwt.NewNumericDate(exp),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}

	token_ptr := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := token_ptr.SignedString(key)

	if err != nil {
		msg := "error while generatin refresh token: " + err.Error()
		return "", exceptions.NewAuthError(500, msg)
	}

	return token, nil

}
