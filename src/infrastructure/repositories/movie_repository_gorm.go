package repositories

import (
	"context"
	"movie-reservation-system/src/domain/entities"
	"movie-reservation-system/src/infrastructure/database"

	"gorm.io/gorm"
)

type MovieRepositoryGorm struct {
	connection *database.Connection
}

func NewMovieRepositoryGorm(connection *database.Connection) *MovieRepositoryGorm {
	return &MovieRepositoryGorm{connection}
}

func (repository *MovieRepositoryGorm) Save(ctx context.Context, movie entities.Movie) error {
	db := repository.connection.GetDB()
	return gorm.G[entities.Movie](db).Create(ctx, &movie)
}

func (repository *MovieRepositoryGorm) GetMovies(ctx context.Context) ([]entities.Movie, error) {
	db := repository.connection.GetDB()
	return gorm.G[entities.Movie](db).Preload("Genre", nil).Find(ctx)
}

func (repository *MovieRepositoryGorm) GetMovieByTitle(ctx context.Context, title string) (entities.Movie, error) {
	db := repository.connection.GetDB()
	return gorm.G[entities.Movie](db).Where("title = ?", title).First(ctx)
}

func (repository *MovieRepositoryGorm) GetMovieByID(ctx context.Context, id uint) (entities.Movie, error) {
	db := repository.connection.GetDB()
	return gorm.G[entities.Movie](db).Where("id = ?", id).First(ctx)
}

func (repository *MovieRepositoryGorm) Update(ctx context.Context, id uint, name string) (int, error) {
	db := repository.connection.GetDB()
	return gorm.G[entities.Movie](db).Where("id = ?", id).Update(ctx, "name", name)
}

func (repository *MovieRepositoryGorm) UpdateImage(ctx context.Context, id uint, image string) (int, error) {
	db := repository.connection.GetDB()
	return gorm.G[entities.Movie](db).Where("id = ?", id).Update(ctx, "image", image)
}

func (repository *MovieRepositoryGorm) Delete(ctx context.Context, id uint) (int, error) {
	db := repository.connection.GetDB()
	return gorm.G[entities.Movie](db).Where("id = ?", id).Delete(ctx)
}
