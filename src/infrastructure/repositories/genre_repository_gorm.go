package repositories

import (
	"context"
	"movie-reservation-system/src/domain/entities"
	"movie-reservation-system/src/infrastructure/database"

	"gorm.io/gorm"
)

type GenreRepositoryGorm struct {
	connection *database.Connection
}

func NewGenreRepositoryGorm(connection *database.Connection) *GenreRepositoryGorm {
	return &GenreRepositoryGorm{connection}
}

func (repository *GenreRepositoryGorm) Save(ctx context.Context, genre entities.Genre) error {
	db := repository.connection.GetDB()
	return gorm.G[entities.Genre](db).Create(ctx, &genre)
}

func (repository *GenreRepositoryGorm) GetGenres(ctx context.Context) ([]entities.Genre, error) {
	db := repository.connection.GetDB()
	return gorm.G[entities.Genre](db).Find(ctx)
}

func (repository *GenreRepositoryGorm) GetGenreByName(ctx context.Context, name string) (entities.Genre, error) {
	db := repository.connection.GetDB()
	return gorm.G[entities.Genre](db).Where("name = ?", name).First(ctx)
}

func (repository *GenreRepositoryGorm) GetGenreByID(ctx context.Context, id uint) (entities.Genre, error) {
	db := repository.connection.GetDB()
	return gorm.G[entities.Genre](db).Where("id = ?", id).First(ctx)
}

func (repository *GenreRepositoryGorm) Update(ctx context.Context, id uint, name string) (int, error) {
	db := repository.connection.GetDB()
	return gorm.G[entities.Genre](db).Where("id = ?", id).Update(ctx, "name", name)
}

func (repository *GenreRepositoryGorm) Delete(ctx context.Context, id uint) (int, error) {
	db := repository.connection.GetDB()
	return gorm.G[entities.Genre](db).Where("id = ?", id).Delete(ctx)
}
