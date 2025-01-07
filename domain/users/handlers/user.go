package handlers

import (
	"context"

	"user-management/domain/users"
)

type UserService interface {
	CreateUser(ctx context.Context, user *users.User) (int, error)
	UpdateUser(ctx context.Context, user *users.User) error
	DeleteUser(ctx context.Context, id int) error
	GetUserByID(ctx context.Context, id int) (*users.User, error)
	ListUsers(ctx context.Context) ([]*users.User, error)
}

func CreateHandler(s UserService) func(ctx context.Context, req users.CreateRequest) (users.CreateResponse, error) {
	return func(ctx context.Context, req users.CreateRequest) (users.CreateResponse, error) {
		user := &users.User{
			Name:  req.Name,
			Email: req.Email,
		}
		id, err := s.CreateUser(ctx, user)
		return users.CreateResponse{Id: id}, err
	}
}

func UpdateHandler(s UserService) func(ctx context.Context, req users.UpdateRequest) (users.UpdateResponse, error) {
	return func(ctx context.Context, req users.UpdateRequest) (users.UpdateResponse, error) {
		user := &users.User{
			ID:    req.ID,
			Name:  req.Name,
			Email: req.Email,
		}
		err := s.UpdateUser(ctx, user)
		return users.UpdateResponse{}, err
	}
}

func DeleteHandler(s UserService) func(ctx context.Context, req users.DeleteRequest) (users.DeleteResponse, error) {
	return func(ctx context.Context, req users.DeleteRequest) (users.DeleteResponse, error) {
		err := s.DeleteUser(ctx, req.ID)
		return users.DeleteResponse{}, err
	}
}

func GetHandler(s UserService) func(ctx context.Context, req users.GetRequest) (users.GetResponse, error) {
	return func(ctx context.Context, req users.GetRequest) (users.GetResponse, error) {
		user, err := s.GetUserByID(ctx, req.ID)
		return users.GetResponse{User: *user}, err
	}
}

func ListHandler(s UserService) func(ctx context.Context, req users.ListRequest) (users.ListResponse, error) {
	return func(ctx context.Context, req users.ListRequest) (users.ListResponse, error) {
		us, err := s.ListUsers(ctx)
		return users.ListResponse{us}, err
	}
}
