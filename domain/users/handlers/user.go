package handlers

import (
	"context"
	"fmt"

	"user-management/domain/users"
)

type UserService interface {
	CreateUser(ctx context.Context, user *users.User) (int, error)
	UpdateUser(ctx context.Context, user *users.User) error
	DeleteUser(ctx context.Context, id int) error
	GetUserByID(ctx context.Context, id int) (*users.User, error)
	ListUsers(ctx context.Context) ([]users.User, error)
}

func CreateHandler(s UserService) func(ctx context.Context, req users.CreateRequest) (users.CreateResponse, error) {
	return func(ctx context.Context, req users.CreateRequest) (users.CreateResponse, error) {
		user := &users.User{
			Name:  req.Name,
			Email: req.Email,
		}
		id, err := s.CreateUser(ctx, user)
		if err != nil {
			return users.CreateResponse{}, fmt.Errorf("s.CreateUser: %v", err)
		}

		return users.CreateResponse{Id: id}, nil
	}
}

func UpdateHandler(s UserService) func(ctx context.Context, req users.UpdateRequest) (users.UpdateResponse, error) {
	return func(ctx context.Context, req users.UpdateRequest) (users.UpdateResponse, error) {
		user := &users.User{
			ID:    req.ID,
			Name:  req.Name,
			Email: req.Email,
		}
		if err := s.UpdateUser(ctx, user); err != nil {
			return users.UpdateResponse{}, fmt.Errorf("s.UpdateUser: %v", err)
		}

		return users.UpdateResponse{}, nil
	}
}

func DeleteHandler(s UserService) func(ctx context.Context, req users.DeleteRequest) (users.DeleteResponse, error) {
	return func(ctx context.Context, req users.DeleteRequest) (users.DeleteResponse, error) {
		if err := s.DeleteUser(ctx, req.ID); err != nil {
			return users.DeleteResponse{}, fmt.Errorf("s.DeleteUser: %v", err)
		}

		return users.DeleteResponse{}, nil
	}
}

func GetHandler(s UserService) func(ctx context.Context, req users.GetRequest) (users.GetResponse, error) {
	return func(ctx context.Context, req users.GetRequest) (users.GetResponse, error) {
		user, err := s.GetUserByID(ctx, req.ID)
		if err != nil {
			return users.GetResponse{}, fmt.Errorf("s.GetUserByID: %v", err)
		}

		return users.GetResponse{User: user}, nil
	}
}

func ListHandler(s UserService) func(ctx context.Context, req users.ListRequest) (users.ListResponse, error) {
	return func(ctx context.Context, req users.ListRequest) (users.ListResponse, error) {
		us, err := s.ListUsers(ctx)
		if err != nil {
			return users.ListResponse{}, fmt.Errorf("s.ListUsers: %v", err)
		}

		return users.ListResponse{Users: us}, nil
	}
}
