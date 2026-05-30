package dto

type UserDTO struct {
	Name     string `json:"name" validate:"required,alphaspace,min=5,max=50"`
	UserName string `json:"username" validate:"required,alphanum,min=5,max=30"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,password"`
}
