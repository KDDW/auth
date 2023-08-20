package dtos

type LoginRefreshTokenDto struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}
