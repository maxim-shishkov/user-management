package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
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
	var id int
	err := r.db.QueryRow(ctx, `INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id`, user.Name, user.Email).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("failed to create user: %w", err)
	}

	return id, nil
}

func (r *UserRepository) Update(ctx context.Context, user *users.User) error {
	_, err := r.db.Exec(ctx, `UPDATE users SET name = $1, email = $2 WHERE id = $3`,
		user.Name, user.Email, user.ID)
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}

	return nil
}

func (r *UserRepository) Delete(ctx context.Context, id int) error {
	_, err := r.db.Exec(ctx, `DELETE FROM users WHERE id = $1`, id)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}

	return nil
}

func (r *UserRepository) GetByID(ctx context.Context, id int) (*users.User, error) {
	row, err := r.db.Query(ctx, `SELECT id, name, email FROM users WHERE id = $1`, id)
	if err != nil {
		return nil, fmt.Errorf("db.Query: %w", err)
	}
	user, err := pgx.CollectExactlyOneRow(row, pgx.RowToStructByName[users.User])
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get user by ID: %w", err)
	}

	return &user, nil
}

func (r *UserRepository) GetAll(ctx context.Context) ([]users.User, error) {
	rows, err := r.db.Query(ctx, `SELECT id, name, email FROM users`)
	if err != nil {
		return nil, fmt.Errorf("db.Query: %w", err)
	}
	us, err := pgx.CollectRows(rows, pgx.RowToStructByName[users.User])
	if err != nil {
		return nil, fmt.Errorf("db.Query: %w", err)
	}

	return us, nil
}
