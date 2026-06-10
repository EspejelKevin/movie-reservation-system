package repositories

import (
	"context"
	"movie-reservation-system/src/domain/dto"
	"movie-reservation-system/src/domain/entities"
	"movie-reservation-system/src/infrastructure/database"

	"gorm.io/gorm"
)

type ShowTimeRepositoryGorm struct {
	connection *database.Connection
}

func NewShowTimeRepositoryGorm(connection *database.Connection) *ShowTimeRepositoryGorm {
	return &ShowTimeRepositoryGorm{connection}
}

func (repository *ShowTimeRepositoryGorm) Save(ctx context.Context, show entities.ShowTime) error {
	db := repository.connection.GetDB()
	return gorm.G[entities.ShowTime](db).Create(ctx, &show)
}

func (repository *ShowTimeRepositoryGorm) GetShows(ctx context.Context) ([]entities.ShowTime, error) {
	db := repository.connection.GetDB()
	return gorm.G[entities.ShowTime](db).Preload("Movie", nil).Find(ctx)
}

func (repository *ShowTimeRepositoryGorm) GetShowByID(ctx context.Context, id uint) (entities.ShowTime, error) {
	db := repository.connection.GetDB()
	return gorm.G[entities.ShowTime](db).Preload("Movie", nil).Where("id = ?", id).First(ctx)
}

func (repository *ShowTimeRepositoryGorm) GetShowByFolio(ctx context.Context, folio string) (entities.ShowTime, error) {
	db := repository.connection.GetDB()
	return gorm.G[entities.ShowTime](db).Where("folio = ?", folio).First(ctx)
}

func (repository *ShowTimeRepositoryGorm) UpdateFolio(ctx context.Context, id uint, folio string) (int, error) {
	db := repository.connection.GetDB()
	return gorm.G[entities.ShowTime](db).Where("id = ?", id).Update(ctx, "folio", folio)
}

func (repository *ShowTimeRepositoryGorm) UpdateUnavailableSeats(ctx context.Context, id uint, seats int) (int, error) {
	db := repository.connection.GetDB()
	return gorm.G[entities.ShowTime](db).Where("id = ?", id).Update(ctx, "UnavailableSeats", seats)
}

func (repository *ShowTimeRepositoryGorm) Update(ctx context.Context, id uint, show dto.ShowDTO) (int, error) {
	db := repository.connection.GetDB()
	return gorm.G[entities.ShowTime](db).Where("id = ?", id).Updates(ctx, entities.ShowTime{StartDate: show.StartDate,
		EndDate: show.EndDate, AvailableQuantitySeats: show.AvailableQuantitySeats, MovieID: show.MovieID})
}

func (repository *ShowTimeRepositoryGorm) Delete(ctx context.Context, id uint) (int, error) {
	db := repository.connection.GetDB()
	return gorm.G[entities.ShowTime](db).Where("id = ?", id).Delete(ctx)
}
