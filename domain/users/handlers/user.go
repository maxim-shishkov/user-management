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

type CreateRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
type CreateResponse struct {
	Id int `json:"id"`
}

func CreateHandler(s UserService) func(ctx context.Context, req CreateRequest) (CreateResponse, error) {
	return func(ctx context.Context, req CreateRequest) (CreateResponse, error) {
		user := &users.User{
			Name:  req.Name,
			Email: req.Email,
		}
		id, err := s.CreateUser(ctx, user)
		return CreateResponse{Id: id}, err
	}
}

type UpdateRequest struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
type UpdateResponse struct{}

func UpdateHandler(s UserService) func(ctx context.Context, req UpdateRequest) (UpdateResponse, error) {
	return func(ctx context.Context, req UpdateRequest) (UpdateResponse, error) {
		user := &users.User{
			ID:    req.ID,
			Name:  req.Name,
			Email: req.Email,
		}
		err := s.UpdateUser(ctx, user)
		return UpdateResponse{}, err
	}
}

type DeleteRequest struct {
	ID int `json:"id"`
}
type DeleteResponse struct{}

func DeleteHandler(s UserService) func(ctx context.Context, req DeleteRequest) (DeleteResponse, error) {
	return func(ctx context.Context, req DeleteRequest) (DeleteResponse, error) {
		err := s.DeleteUser(ctx, req.ID)
		return DeleteResponse{}, err
	}
}

type GetRequest struct {
	ID int `json:"id"`
}
type GetResponse struct {
	User users.User `json:"user"`
}

func GetHandler(s UserService) func(ctx context.Context, req GetRequest) (GetResponse, error) {
	return func(ctx context.Context, req GetRequest) (GetResponse, error) {
		user, err := s.GetUserByID(ctx, req.ID)
		return GetResponse{User: *user}, err
	}
}

type ListRequest struct{}
type ListResponse struct {
	Users []*users.User `json:"users"`
}

func ListHandler(s UserService) func(ctx context.Context, req ListRequest) (ListResponse, error) {
	return func(ctx context.Context, req ListRequest) (ListResponse, error) {
		users, err := s.ListUsers(ctx)
		return ListResponse{users}, err
	}
}
