package interfaces

import (
	"context"
	"movie-reservation-system/src/domain/entities"
)

type RoleRepository interface {
	Save(ctx context.Context, role entities.Role) error
	GetRoleByName(ctx context.Context, name string) (entities.Role, error)
	GetRoleByID(ctx context.Context, id uint) (entities.Role, error)
	Update(ctx context.Context, id uint, name string) (int, error)
	Delete(ctx context.Context, id uint) (int, error)
}
