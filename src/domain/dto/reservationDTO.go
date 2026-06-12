package dto

type ReservationDTO struct {
	Folio         string `json:"folio" validate:"required,alphanum,min=8,max=12"`
	ReservedSeats int    `json:"reserved_seats" validate:"required,gte=1"`
	ShowTimeID    uint   `json:"showtime_id" validate:"required,gte=1"`
}

type ReservationResponse struct {
	ID       uint         `json:"id"`
	Folio    string       `json:"folio"`
	Status   string       `json:"status"`
	ShowTime ShowResponse `json:"showtime"`
	User     UserResponse `json:"user"`
}
