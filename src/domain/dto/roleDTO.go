package dto

type RoleDTO struct {
	Name string `json:"name" validate:"role"`
}
