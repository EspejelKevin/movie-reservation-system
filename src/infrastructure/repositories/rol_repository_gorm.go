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

func (repository *RolRepositoryGorm) Save(role entities.Role) error {
	ctx := context.Background()
	db := repository.connection.GetDB()
	return gorm.G[entities.Role](db).Create(ctx, &role)
}

func (repository *RolRepositoryGorm) GetRoleByName(name string) (entities.Role, error) {
	ctx := context.Background()
	db := repository.connection.GetDB()
	return gorm.G[entities.Role](db).Where("name = ?", name).First(ctx)
}

func (repository *RolRepositoryGorm) GetRoleByID(id uint) (entities.Role, error) {
	ctx := context.Background()
	db := repository.connection.GetDB()
	return gorm.G[entities.Role](db).Where("id = ?", id).First(ctx)
}

func (repository *RolRepositoryGorm) Update(id uint, name string) (int, error) {
	ctx := context.Background()
	db := repository.connection.GetDB()
	return gorm.G[entities.Role](db).Where("id = ?", id).Update(ctx, "name", name)
}

func (repository *RolRepositoryGorm) Delete(id uint) (int, error) {
	ctx := context.Background()
	db := repository.connection.GetDB()
	return gorm.G[entities.Role](db).Where("id = ?", id).Delete(ctx)
}
