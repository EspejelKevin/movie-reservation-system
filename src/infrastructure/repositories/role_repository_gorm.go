package repositories

import (
	"context"
	"movie-reservation-system/src/domain/entities"
	"movie-reservation-system/src/infrastructure/database"

	"gorm.io/gorm"
)

type RoleRepositoryGorm struct {
	connection *database.Connection
}

func NewRoleRepositoryGorm(connection *database.Connection) *RoleRepositoryGorm {
	return &RoleRepositoryGorm{connection}
}

func (repository *RoleRepositoryGorm) Save(ctx context.Context, role entities.Role) error {
	db := repository.connection.GetDB()
	return gorm.G[entities.Role](db).Create(ctx, &role)
}

func (repository *RoleRepositoryGorm) GetRoles(ctx context.Context) ([]entities.Role, error) {
	db := repository.connection.GetDB()
	return gorm.G[entities.Role](db).Find(ctx)
}

func (repository *RoleRepositoryGorm) GetRoleByName(ctx context.Context, name string) (entities.Role, error) {
	db := repository.connection.GetDB()
	return gorm.G[entities.Role](db).Where("name = ?", name).First(ctx)
}

func (repository *RoleRepositoryGorm) GetRoleByID(ctx context.Context, id uint) (entities.Role, error) {
	db := repository.connection.GetDB()
	return gorm.G[entities.Role](db).Where("id = ?", id).First(ctx)
}

func (repository *RoleRepositoryGorm) Update(ctx context.Context, id uint, name string) (int, error) {
	db := repository.connection.GetDB()
	return gorm.G[entities.Role](db).Where("id = ?", id).Update(ctx, "name", name)
}

func (repository *RoleRepositoryGorm) Delete(ctx context.Context, id uint) (int, error) {
	db := repository.connection.GetDB()
	return gorm.G[entities.Role](db).Where("id = ?", id).Delete(ctx)
}
