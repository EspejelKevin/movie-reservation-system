package repositories

import (
	"context"
	"movie-reservation-system/src/domain/entities"
	"movie-reservation-system/src/infrastructure/database"

	"gorm.io/gorm"
)

type ReservationRepositoryGorm struct {
	connection *database.Connection
}

func NewReservationRepositoryGorm(connection *database.Connection) *ReservationRepositoryGorm {
	return &ReservationRepositoryGorm{connection}
}

func (repository *ReservationRepositoryGorm) Save(ctx context.Context, reservation entities.Reservation) error {
	db := repository.connection.GetDB()
	return gorm.G[entities.Reservation](db).Create(ctx, &reservation)
}

func (repository *ReservationRepositoryGorm) GetReservations(ctx context.Context) ([]entities.Reservation, error) {
	db := repository.connection.GetDB()
	return gorm.G[entities.Reservation](db).
		Preload("ShowTime.Movie.Genre", nil).
		Preload("User", nil).Find(ctx)
}

func (repository *ReservationRepositoryGorm) GetReservationByID(ctx context.Context, id uint) (entities.Reservation, error) {
	db := repository.connection.GetDB()
	return gorm.G[entities.Reservation](db).
		Preload("ShowTime.Movie.Genre", nil).
		Preload("User", nil).
		Where("id = ?", id).First(ctx)
}

func (repository *ReservationRepositoryGorm) GetReservationByFolio(ctx context.Context, folio string) (entities.Reservation, error) {
	db := repository.connection.GetDB()
	return gorm.G[entities.Reservation](db).
		Preload("ShowTime.Movie.Genre", nil).
		Preload("User", nil).
		Where("folio = ?", folio).First(ctx)
}

func (repository *ReservationRepositoryGorm) GetReservationsByUserID(ctx context.Context, userID uint) ([]entities.Reservation, error) {
	db := repository.connection.GetDB()
	return gorm.G[entities.Reservation](db).
		Preload("ShowTime.Movie.Genre", nil).
		Preload("User", nil).
		Where("user_id = ?", userID).Find(ctx)
}

func (repository *ReservationRepositoryGorm) GetReservationByIDAndUserID(ctx context.Context, id uint, userID uint) (entities.Reservation, error) {
	db := repository.connection.GetDB()
	return gorm.G[entities.Reservation](db).
		Preload("ShowTime.Movie.Genre", nil).
		Preload("User", nil).
		Where("id = ? and user_id = ?", id, userID).First(ctx)
}

func (repository *ReservationRepositoryGorm) UpdateReservationStatus(ctx context.Context, id uint, userID uint, status string) (int, error) {
	db := repository.connection.GetDB()
	return gorm.G[entities.Reservation](db).Where("id = ? and user_id = ?", id, userID).Update(ctx, "status", status)
}

func (repository *ReservationRepositoryGorm) Delete(ctx context.Context, id uint) (int, error) {
	db := repository.connection.GetDB()
	return gorm.G[entities.Reservation](db).Where("id = ?", id).Delete(ctx)
}
