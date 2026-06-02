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

func (repository *UserRepositoryGorm) Save(ctx context.Context, user entities.User) error {
	db := repository.connection.GetDB()
	return gorm.G[entities.User](db).Create(ctx, &user)
}

func (repository *UserRepositoryGorm) GetUserByEmailOrUserName(ctx context.Context, email string,
	username string) (entities.User, error) {

	db := repository.connection.GetDB()
	return gorm.G[entities.User](db).Where("email = ?", email).Or("username = ?", username).First(ctx)
}
