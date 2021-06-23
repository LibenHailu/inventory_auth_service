package userservice

import (
	"context"

	"github.com/LibenHailu/inventory_auth/internal/core/domain"
	"github.com/LibenHailu/inventory_auth/internal/core/ports"
)

type userService struct {
	userRepo ports.UserRepository
}

// creates a user Service object
func NewUserService(userRepo ports.UserRepository) *userService {
	return &userService{
		userRepo: userRepo,
	}
}

func (ur userService) StoreUser(ctx context.Context, user *domain.User) (*domain.User, error) {

	user, err := ur.userRepo.StoreUser(ctx, user)

	if err != nil {
		return nil, err
	}
	return user, nil
}
