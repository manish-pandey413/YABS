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

func (u *UserRepository) AddUser(ctx context.Context, username string, email string, password string) (*model.User, error) {
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

	rows, err := u.server.DB.Pool.Query(ctx, stmt, pgx.NamedArgs{
		"username": username,
		"email":    email,
		"password": password,
	})

	if err != nil {
		return nil, fmt.Errorf("Failed to add user %w\n", err)
	}

	userItem, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[model.User])
	if err != nil {
		return nil, fmt.Errorf("Failed to collect user. %w \n", err)
	}

	return &userItem, nil
}

func (u *UserRepository) GetUser(ctx context.Context, username string) (*model.User, error) {
	stmt := `
		SELECT 
			*
		FROM 
			users
		WHERE
			username = @username
	`

	rows, err := u.server.DB.Pool.Query(ctx, stmt, pgx.NamedArgs{
		"username": username,
	})

	if err != nil {
		return nil, fmt.Errorf("failed to execute select query for username= %s :%w", username, err)
	}

	user, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[model.User])
	if err != nil {
		return nil, fmt.Errorf("failed to collect row from table:users for username=%s : %w", username, err)
	}
	return &user, nil
}
