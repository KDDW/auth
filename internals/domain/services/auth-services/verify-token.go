package auth_services

import (
	"auth-service/internals/utils/exceptions"
	"auth-service/internals/utils/tokens"
	"strconv"
)

func (a *AuthServices) VerifyToken(token string) (bool, *exceptions.AuthError) {

	claims, err := tokens.VerifyToken(token)

	if err != nil {
		return false, exceptions.NewAuthError(401, "invalid authorization: "+err.Error())
	}

	userId, err := strconv.ParseInt(claims.Subject, 10, 64)

	if err != nil {
		authError := exceptions.NewAuthError(401, "invalid authorization: "+err.Error())
		return false, authError
	}

	_, authErr := a.UserServices.GetById(userId)

	if authErr != nil {
		return false, authErr
	}

	return true, nil
}
