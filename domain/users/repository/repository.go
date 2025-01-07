package repository

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"user-management/domain/users"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(dsn string) (*UserRepository, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}
	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("ping %v", err)
	}

	return &UserRepository{db: db}, nil
}

func (r *UserRepository) Create(ctx context.Context, user *users.User) (int, error) {
	query := `INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id`
	var id int
	err := r.db.QueryRowContext(ctx, query, user.Name, user.Email).Scan(&id)
	return id, err
}

func (r *UserRepository) Update(ctx context.Context, user *users.User) error {
	query := `UPDATE users SET name = $1, email = $2  WHERE id = $3`
	_, err := r.db.ExecContext(ctx, query, user.Name, user.Email, user.ID)
	return err
}

func (r *UserRepository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

func (r *UserRepository) GetByID(ctx context.Context, id int) (*users.User, error) {
	query := `SELECT id, name, email FROM users WHERE id = $1`
	user := &users.User{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(&user.ID, &user.Name, &user.Email)
	return user, err
}

func (r *UserRepository) GetAll(ctx context.Context) ([]*users.User, error) {
	query := `SELECT id, name, email FROM users`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var us []*users.User
	for rows.Next() {
		user := &users.User{}
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			return nil, err
		}
		us = append(us, user)
	}

	return us, nil
}
