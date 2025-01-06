package domain

import "context"

type User struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRepository interface {
	Create(ctx context.Context, user *User) (int, error)
	Update(ctx context.Context, user *User) error
	Delete(ctx context.Context, id int) error
	GetByID(ctx context.Context, id int) (*User, error)
	GetAll(ctx context.Context) ([]*User, error)
}

//
//type UserService interface {
//	CreateUser(ctx context.Context, user *User) (int64, error)
//	UpdateUser(ctx context.Context, user *User) error
//	DeleteUser(ctx context.Context, id int) error
//	GetUserByID(ctx context.Context, id int) (*User, error)
//	ListUsers(ctx context.Context) ([]*User, error)
//}
