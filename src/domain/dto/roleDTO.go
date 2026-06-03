package dto

type RoleDTO struct {
	Name string `json:"name" validate:"role,min=5,max=30"`
}
