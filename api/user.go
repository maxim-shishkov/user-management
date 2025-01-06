package api

import (
	"context"

	"user-management/domain"
)

type UserService interface {
	CreateUser(ctx context.Context, user *domain.User) (int, error)
	UpdateUser(ctx context.Context, user *domain.User) error
	DeleteUser(ctx context.Context, id int) error
	GetUserByID(ctx context.Context, id int) (*domain.User, error)
	ListUsers(ctx context.Context) ([]*domain.User, error)
}

type CreateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func CreateUserHandler(s UserService) func(ctx context.Context, req CreateUserRequest) (any, error) {
	return func(ctx context.Context, req CreateUserRequest) (any, error) {
		user := &domain.User{
			Name:  req.Name,
			Email: req.Email,
		}
		return s.CreateUser(ctx, user)
	}
}

type UpdateUserRequest struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func UpdateUserHandler(s UserService) func(ctx context.Context, req UpdateUserRequest) (any, error) {
	return func(ctx context.Context, req UpdateUserRequest) (any, error) {
		user := &domain.User{
			ID:    req.ID,
			Name:  req.Name,
			Email: req.Email,
		}
		return nil, s.UpdateUser(ctx, user)
	}
}

type DeleteUserRequest struct {
	ID int `json:"id"`
}

func DeleteUserHandler(s UserService) func(ctx context.Context, req DeleteUserRequest) (any, error) {
	return func(ctx context.Context, req DeleteUserRequest) (any, error) {
		return nil, s.DeleteUser(ctx, req.ID)
	}
}

type GetUserRequest struct {
	ID int `json:"id"`
}

func GetUserHandler(s UserService) func(ctx context.Context, req GetUserRequest) (any, error) {
	return func(ctx context.Context, req GetUserRequest) (any, error) {
		return s.GetUserByID(ctx, req.ID)
	}
}

type ListUserRequest struct{}

func ListUsersHandler(s UserService) func(ctx context.Context, req ListUserRequest) (any, error) {
	return func(ctx context.Context, req ListUserRequest) (any, error) {
		return s.ListUsers(ctx)
	}
}
