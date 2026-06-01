package interfaces

import "movie-reservation-system/src/domain/entities"

type RoleRepository interface {
	Save(role entities.Role) error
	GetRoleByName(name string) (entities.Role, error)
	GetRoleByID(id uint) (entities.Role, error)
	Update(id uint, name string) (int, error)
	Delete(id uint) (int, error)
}
