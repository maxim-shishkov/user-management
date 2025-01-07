package service

import (
	"context"
	"fmt"

	"go.uber.org/zap"
	"user-management/domain/users"
)

type UserRepository interface {
	Create(ctx context.Context, user *users.User) (int, error)
	Update(ctx context.Context, user *users.User) error
	Delete(ctx context.Context, id int) error
	GetByID(ctx context.Context, id int) (*users.User, error)
	GetAll(ctx context.Context) ([]*users.User, error)
}

type UserService struct {
	repo   UserRepository
	logger *zap.Logger
}

func NewUserService(repo UserRepository, logger *zap.Logger) *UserService {
	return &UserService{repo: repo, logger: logger}
}

func (s *UserService) CreateUser(ctx context.Context, user *users.User) (int, error) {
	id, err := s.repo.Create(ctx, user)
	if err != nil {
		return 0, fmt.Errorf("repo.Create: %v", err)
	}

	return id, nil
}

func (s *UserService) UpdateUser(ctx context.Context, user *users.User) error {
	if err := s.repo.Update(ctx, user); err != nil {
		return fmt.Errorf("repo.Update: %v", err)
	}

	return nil
}

func (s *UserService) DeleteUser(ctx context.Context, id int) error {
	if err := s.repo.Delete(ctx, id); err != nil {
		return fmt.Errorf("repo.Delete: %v", err)
	}

	return nil
}

func (s *UserService) GetUserByID(ctx context.Context, id int) (*users.User, error) {
	u, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("repo.GetByID: %v", err)
	}

	return u, nil
}

func (s *UserService) ListUsers(ctx context.Context) ([]*users.User, error) {
	us, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("repo.GetAll: %v", err)
	}

	return us, nil
}
