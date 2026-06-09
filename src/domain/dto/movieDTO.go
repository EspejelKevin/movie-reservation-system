package dto

type MovieDTO struct {
	Title       string `json:"title" validate:"required,min=5,max=30"`
	Description string `json:"description" validate:"required,min=5,max=100"`
	GenreID     uint   `json:"genre_id" validate:"required,gte=1"`
}

type MovieResponse struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Genre       string `json:"genre"`
}
