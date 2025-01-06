package service

import (
	"context"
	"fmt"

	"user-management/domain"
)

type UserService struct {
	repo domain.UserRepository
}

func NewUserService(repo domain.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(ctx context.Context, req CreateUserRequest) (any, error) {
	return UserResponse{}, nil
}

func (s *UserService) GetUser(ctx context.Context, req GetUserRequest) (any, error) {

	if req.ID != 12345 {
		return nil, fmt.Errorf("qwe error")
	}

	user, err := s.repo.GetByID(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return UserResponse{
		ID:    fmt.Sprint(user.ID),
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func (s *UserService) UpdateUser(ctx context.Context, req UpdateUserRequest) (any, error) {
	return UserResponse{}, nil
}

func (s *UserService) DeleteUser(ctx context.Context, req DeleteUserRequest) (any, error) {
	return nil, nil
}
