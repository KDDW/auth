package dtos

type CreateUserDto struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
	Realm    string `json:"realm" validate:"required"`
	Name     string `json:"name" validate:"required"`
}
