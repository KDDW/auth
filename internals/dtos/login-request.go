package dtos

type LoginRequestDto struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
	Realm    string `json:"realm" validate:"required"`
}
