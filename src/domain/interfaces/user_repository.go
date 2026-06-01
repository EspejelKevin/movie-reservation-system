package interfaces

import "movie-reservation-system/src/domain/entities"

type UserRepository interface {
	Save(user entities.User) error
	GetUserByEmail(email string) (entities.User, error)
}
