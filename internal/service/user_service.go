package service

import (
	"context"
	"github.com/DDnagasawa/FocusFlow/internal/model"
	"github.com/DDnagasawa/FocusFlow/internal/repository"
)

// UserService handles user-related business logic
type UserService struct {
	repo repository.UserRepository
}

// NewUserService creates a new UserService
func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

// CreateUser creates a new user
func (s *UserService) CreateUser(ctx context.Context, user *model.User) error {
	// Add any business logic here (e.g., validation)
	return s.repo.CreateUser(ctx, user)
}

// GetUser retrieves a user by ID
func (s *UserService) GetUser(ctx context.Context, id int) (*model.User, error) {
	return s.repo.GetUser(ctx, id)
}

// UpdateUser updates a user's information
func (s *UserService) UpdateUser(ctx context.Context, user *model.User) error {
	// Add any business logic here (e.g., validation)
	return s.repo.UpdateUser(ctx, user)
}

// DeleteUser deletes a user
func (s *UserService) DeleteUser(ctx context.Context, id int) error {
	return s.repo.DeleteUser(ctx, id)
}
