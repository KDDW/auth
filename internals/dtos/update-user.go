package dtos

type UpdateUserDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
