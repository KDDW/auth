package dtos

type CreateUserDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Realm    string `json:"realm"`
}
