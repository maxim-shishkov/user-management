package service

import (
	"context"

	"user-management/domain"
)

type UserRepository interface {
	Create(ctx context.Context, user *domain.User) (int, error)
	Update(ctx context.Context, user *domain.User) error
	Delete(ctx context.Context, id int) error
	GetByID(ctx context.Context, id int) (*domain.User, error)
	GetAll(ctx context.Context) ([]*domain.User, error)
}

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(ctx context.Context, user *domain.User) (int, error) {
	return s.repo.Create(ctx, user)
}

func (s *UserService) UpdateUser(ctx context.Context, user *domain.User) error {
	return s.repo.Update(ctx, user)
}

func (s *UserService) DeleteUser(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}

func (s *UserService) GetUserByID(ctx context.Context, id int) (*domain.User, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *UserService) ListUsers(ctx context.Context) ([]*domain.User, error) {
	return s.repo.GetAll(ctx)
}
