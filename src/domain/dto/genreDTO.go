package dto

type GenreDTO struct {
	Name string `json:"name" validate:"min=5,max=30"`
}
