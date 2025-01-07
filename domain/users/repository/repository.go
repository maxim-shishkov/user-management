package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"user-management/domain/users"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(dsn string) (*UserRepository, error) {
	db, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}
	return &UserRepository{db: db}, nil
}

func (r *UserRepository) Close() {
	r.db.Close()
}

func (r *UserRepository) Create(ctx context.Context, user *users.User) (int, error) {
	query := `INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id`
	var id int
	err := r.db.QueryRow(ctx, query, user.Name, user.Email).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("failed to create user: %w", err)
	}
	return id, nil
}

func (r *UserRepository) Update(ctx context.Context, user *users.User) error {
	query := `UPDATE users SET name = $1, email = $2 WHERE id = $3`
	_, err := r.db.Exec(ctx, query, user.Name, user.Email, user.ID)
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}
	return nil
}

func (r *UserRepository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}
	return nil
}

func (r *UserRepository) GetByID(ctx context.Context, id int) (*users.User, error) {
	query := `SELECT id, name, email FROM users WHERE id = $1`
	user := &users.User{}
	err := r.db.QueryRow(ctx, query, id).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by ID: %w", err)
	}
	return user, nil
}

func (r *UserRepository) GetAll(ctx context.Context) ([]*users.User, error) {
	query := `SELECT id, name, email FROM users`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get all users: %w", err)
	}
	defer rows.Close()

	var usersList []*users.User
	for rows.Next() {
		user := &users.User{}
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			return nil, fmt.Errorf("failed to scan user: %w", err)
		}
		usersList = append(usersList, user)
	}
	if rows.Err() != nil {
		return nil, fmt.Errorf("rows error: %w", rows.Err())
	}
	return usersList, nil
}
