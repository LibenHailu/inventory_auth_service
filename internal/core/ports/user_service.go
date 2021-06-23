package ports

import (
	"context"

	"github.com/LibenHailu/inventory_auth/internal/core/domain"
)

type UserService interface {
	// Users() ([]domain.User, []error)
	// User(id uint) (*domain.User, []error)
	StoreUser(ctx context.Context, user *domain.User) (*domain.User, error)
	// UserByEmail(email string) (*domain.User, []error)
	// UpdateUser(user *domain.User) (*domain.User, []error)
	// DeleteUser(id uint) (*domain.User, []error)
}
