package service

import (
	"context"

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
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(ctx context.Context, user *users.User) (int, error) {
	return s.repo.Create(ctx, user)
}

func (s *UserService) UpdateUser(ctx context.Context, user *users.User) error {
	return s.repo.Update(ctx, user)
}

func (s *UserService) DeleteUser(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}

func (s *UserService) GetUserByID(ctx context.Context, id int) (*users.User, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *UserService) ListUsers(ctx context.Context) ([]*users.User, error) {
	return s.repo.GetAll(ctx)
}
