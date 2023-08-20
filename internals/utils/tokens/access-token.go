package tokens

import (
	"auth-service/internals/utils/exceptions"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateAccessToken(userId int64) (string, *exceptions.AuthError) {

	key := []byte(os.Getenv("JWT_SECRET"))

	exp := time.Now().UTC().Add(3 * time.Hour) // 3 hours

	claims := jwt.RegisteredClaims{
		Issuer:    "kddw-auth-service",
		Subject:   strconv.FormatInt(userId, 10),
		ExpiresAt: jwt.NewNumericDate(exp),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}

	token_ptr := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := token_ptr.SignedString(key)

	if err != nil {
		msg := "error while generatin access token: " + err.Error()
		fmt.Println(msg)
		return "", exceptions.NewAuthError(500, "unknow error")
	}

	return token, nil

}
