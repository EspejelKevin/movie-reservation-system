package interfaces

import (
	"context"
	"movie-reservation-system/src/domain/entities"
)

type GenreRepository interface {
	Save(ctx context.Context, genre entities.Genre) error
	GetGenres(ctx context.Context) ([]entities.Genre, error)
	GetGenreByName(ctx context.Context, name string) (entities.Genre, error)
	GetGenreByID(ctx context.Context, id uint) (entities.Genre, error)
	Update(ctx context.Context, id uint, name string) (int, error)
	Delete(ctx context.Context, id uint) (int, error)
}
