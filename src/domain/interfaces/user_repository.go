package interfaces

import (
	"context"
	"movie-reservation-system/src/domain/entities"
)

type UserRepository interface {
	Save(ctx context.Context, user entities.User) error
	GetUserByEmail(ctx context.Context, email string) (entities.User, error)
}
