package repositories

import (
	"context"
	"movie-reservation-system/src/domain/entities"
	"movie-reservation-system/src/infrastructure/database"

	"gorm.io/gorm"
)

type UserRepositoryGorm struct {
	connection *database.Connection
}

func NewUserRepositoryGorm(connection *database.Connection) *UserRepositoryGorm {
	return &UserRepositoryGorm{connection}
}

func (repository *UserRepositoryGorm) Save(user entities.User) error {
	ctx := context.Background()
	db := repository.connection.GetDB()
	return gorm.G[entities.User](db).Create(ctx, &user)
}

func (repository *UserRepositoryGorm) GetUserByEmail(email string) (entities.User, error) {
	ctx := context.Background()
	db := repository.connection.GetDB()
	return gorm.G[entities.User](db).Where("email = ?", email).First(ctx)
}
