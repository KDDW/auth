package auth_services

import (
	"auth-service/internals/dtos"
	"auth-service/internals/utils/exceptions"
	"auth-service/internals/utils/tokens"
	"strconv"
)

func (a *AuthServices) LoginRefreshToken(refreshToken string) (*dtos.LoginResponse, *exceptions.AuthError) {

	claims, err := tokens.VerifyToken(refreshToken)

	if err != nil {
		return nil, exceptions.NewAuthError(401, "invalid authorization: "+err.Error())
	}

	userId, err := strconv.ParseInt(claims.Subject, 10, 64)

	if err != nil {
		authErr := exceptions.NewAuthError(401, "invalid authorization: "+err.Error())
		return nil, authErr
	}

	userFound, authErr := a.UserServices.GetById(userId)

	if authErr != nil {
		return nil, authErr
	}

	accessToken, authErr := tokens.GenerateAccessToken(userFound.ID)

	if authErr != nil {
		return nil, authErr
	}

	refresh, authErr := tokens.GenerateRefreshToken(userFound.ID)

	if authErr != nil {
		return nil, authErr
	}

	response := &dtos.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refresh,
	}

	return response, nil
}
