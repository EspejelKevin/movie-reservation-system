package dto

import "time"

type ShowDTO struct {
	StartDate              time.Time `json:"start_date" validate:"required,datetime"`
	EndDate                time.Time `json:"end_date" validate:"required,datetime"`
	AvailableQuantitySeats int       `json:"available_quantity_seats" validate:"required,gte=1"`
	MovieID                uint      `json:"movie_id" validate:"required,gte=1"`
}
