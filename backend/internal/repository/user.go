package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/manish-pandey413/YABS/internal/model"
	"github.com/manish-pandey413/YABS/internal/server"
)

type UserRepository struct {
	server *server.Server
}

func NewUserRepository(s *server.Server) *UserRepository {
	return &UserRepository{
		server: s,
	}
}

func (ur *UserRepository) AddUser(ctx context.Context, username string, email string) (*model.User, error) {
	stmt := `
		INSERT INTO 
			users (
				username,
				email,
				password
			)
			VALUES
			(
				@username,
				@email,
				@password
			)
			RETURNING *
	`

	rows, err := ur.server.DB.Pool.Query(ctx, stmt, pgx.NamedArgs{
		"username": username,
		"email":    email,
		"password": "12345",
	})

	if err != nil {
		return nil, fmt.Errorf("Failed to add user with username %s: %w", username, err)
	}

	userItem, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[model.User])
	if err != nil {
		return nil, fmt.Errorf("Failed to collect user %w", err)
	}

	return &userItem, nil
}
