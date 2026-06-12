package dto

type ShowDTO struct {
	Folio         string `json:"folio" validate:"required,alphanum,min=8,max=12"`
	StartDate     string `json:"start_date" validate:"required,datetime=2006-01-02 15:04:05"`
	EndDate       string `json:"end_date" validate:"required,datetime=2006-01-02 15:04:05,after_start"`
	QuantitySeats int    `json:"available_quantity_seats" validate:"required,gte=1"`
	MovieID       uint   `json:"movie_id" validate:"required,gte=1"`
}

type ShowResponse struct {
	ID                     uint          `json:"id"`
	Folio                  string        `json:"folio"`
	StartDate              string        `json:"start_date"`
	EndDate                string        `json:"end_date"`
	AvailableQuantitySeats int           `json:"available_quantity_seats"`
	UnavailableSeats       *int          `json:"unavailable_seats"`
	Movie                  MovieResponse `json:"movie"`
}
