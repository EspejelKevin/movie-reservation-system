package interfaces

import (
	"context"
	"movie-reservation-system/src/domain/entities"
)

type ReservationRepository interface {
	Save(ctx context.Context, reservation entities.Reservation) error
	GetReservationByID(ctx context.Context, id uint) (entities.Reservation, error)
	GetReservationByFolio(ctx context.Context, folio string) (entities.Reservation, error)
	GetReservationsByUserID(ctx context.Context, userID uint) ([]entities.Reservation, error)
}
