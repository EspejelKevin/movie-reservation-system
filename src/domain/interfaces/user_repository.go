package interfaces

import (
	"context"
	"movie-reservation-system/src/domain/entities"
)

type UserRepository interface {
	Save(ctx context.Context, user entities.User) error
	GetUserByEmailOrUserName(ctx context.Context, email string, username string) (entities.User, error)
}
