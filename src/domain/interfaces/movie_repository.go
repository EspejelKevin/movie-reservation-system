package interfaces

import (
	"context"
	"movie-reservation-system/src/domain/entities"
)

type MovieRepository interface {
	Save(ctx context.Context, movie entities.Movie) error
	GetMovies(ctx context.Context) ([]entities.Movie, error)
	GetMovieByID(ctx context.Context, id uint) (entities.Movie, error)
	GetMovieByTitle(ctx context.Context, title string) (entities.Movie, error)
	Update(ctx context.Context, id uint, name string) (int, error)
	UpdateImage(ctx context.Context, id uint, image string) (int, error)
	Delete(ctx context.Context, id uint) (int, error)
}
