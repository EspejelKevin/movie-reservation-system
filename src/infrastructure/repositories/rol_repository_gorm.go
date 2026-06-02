package repositories

import (
	"context"
	"movie-reservation-system/src/domain/entities"
	"movie-reservation-system/src/infrastructure/database"

	"gorm.io/gorm"
)

type RolRepositoryGorm struct {
	connection *database.Connection
}

func NewRolRepositoryGorm(connection *database.Connection) *RolRepositoryGorm {
	return &RolRepositoryGorm{connection}
}

func (repository *RolRepositoryGorm) Save(ctx context.Context, role entities.Role) error {
	db := repository.connection.GetDB()
	return gorm.G[entities.Role](db).Create(ctx, &role)
}

func (repository *RolRepositoryGorm) GetRoleByName(ctx context.Context, name string) (entities.Role, error) {
	db := repository.connection.GetDB()
	return gorm.G[entities.Role](db).Where("name = ?", name).First(ctx)
}

func (repository *RolRepositoryGorm) GetRoleByID(ctx context.Context, id uint) (entities.Role, error) {
	db := repository.connection.GetDB()
	return gorm.G[entities.Role](db).Where("id = ?", id).First(ctx)
}

func (repository *RolRepositoryGorm) Update(ctx context.Context, id uint, name string) (int, error) {
	db := repository.connection.GetDB()
	return gorm.G[entities.Role](db).Where("id = ?", id).Update(ctx, "name", name)
}

func (repository *RolRepositoryGorm) Delete(ctx context.Context, id uint) (int, error) {
	db := repository.connection.GetDB()
	return gorm.G[entities.Role](db).Where("id = ?", id).Delete(ctx)
}
