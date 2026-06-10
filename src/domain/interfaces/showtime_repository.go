package interfaces

import (
	"context"
	"movie-reservation-system/src/domain/dto"
	"movie-reservation-system/src/domain/entities"
)

type ShowTimeRepository interface {
	Save(ctx context.Context, show entities.ShowTime) error
	GetShows(ctx context.Context) ([]entities.ShowTime, error)
	GetShowByID(ctx context.Context, id uint) (entities.ShowTime, error)
	GetShowByFolio(ctx context.Context, folio string) (entities.ShowTime, error)
	UpdateFolio(ctx context.Context, id uint, folio string) (int, error)
	UpdateUnavailableSeats(ctx context.Context, id uint, seats int) (int, error)
	Update(ctx context.Context, id uint, show dto.ShowDTO) (int, error)
	Delete(ctx context.Context, id uint) (int, error)
}
