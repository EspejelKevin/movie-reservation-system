package interfaces

import "movie-reservation-system/src/domain/entities"

type RoleRepository interface {
	Save(role entities.Role) error
	GetRoleByName(role string) (*entities.Role, error)
	GetRoleByID(id uint) (*entities.Role, error)
	Update(id uint, role entities.Role) error
	Delete(id uint) error
}
