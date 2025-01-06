package repository

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"user-management/domain"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(dsn string) (*UserRepository, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}
	return &UserRepository{db: db}, nil
}

func (r *UserRepository) Create(ctx context.Context, user *domain.User) (int, error) {
	query := `INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id`
	var id int
	err := r.db.QueryRowContext(ctx, query, user.Name, user.Email, user.Password).Scan(&id)
	return id, err
}

func (r *UserRepository) Update(ctx context.Context, user *domain.User) error {
	query := `UPDATE users SET name = $1, email = $2, password = $3 WHERE id = $4`
	_, err := r.db.ExecContext(ctx, query, user.Name, user.Email, user.Password, user.ID)
	return err
}

func (r *UserRepository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

func (r *UserRepository) GetByID(ctx context.Context, id int) (*domain.User, error) {

	return &domain.User{
		ID:       1,
		Name:     "Max",
		Email:    "email@mail.ru",
		Password: "qwe123",
	}, nil

	query := `SELECT id, name, email, password FROM users WHERE id = $1`
	user := &domain.User{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	return user, err
}

func (r *UserRepository) GetAll(ctx context.Context) ([]*domain.User, error) {
	query := `SELECT id, name, email, password FROM users`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*domain.User
	for rows.Next() {
		user := &domain.User{}
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
