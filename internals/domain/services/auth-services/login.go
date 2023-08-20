package auth_services

import (
	"auth-service/internals/dtos"
	"auth-service/internals/utils/exceptions"
	"auth-service/internals/utils/tokens"
)

func (a *AuthServices) Login(email string, password string, realmCode string) (*dtos.LoginResponse, *exceptions.AuthError) {

	userFound, err := a.UserServices.GetByEmailAndRealm(email, realmCode)

	if err != nil {
		if err.Details == "user not found" {
			return nil, exceptions.NewAuthError(401, "wrong email or password")
		}
		return nil, err
	}

	passwordIsValid := userFound.ComparePassword(password)

	if !passwordIsValid {
		return nil, exceptions.NewAuthError(401, "wrong email or password")
	}

	accessToken, err := tokens.GenerateAccessToken(userFound.ID)

	if err != nil {
		return nil, err
	}

	refreshToken, err := tokens.GenerateRefreshToken(userFound.ID)

	if err != nil {
		return nil, err
	}

	response := &dtos.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return response, nil
}
