package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/manish-pandey413/YABS/internal/model"
	"github.com/manish-pandey413/YABS/internal/server"
)

type PostRepository struct {
	server *server.Server
}

func NewPostRepository(s *server.Server) *PostRepository {
	return &PostRepository{
		server: s,
	}
}

func (p *PostRepository) NewPost(ctx context.Context, owner_username string, content string) (*model.Post, error) {
	stmt := `
		INSERT INTO 
			posts (
				owner_username,
				content,
				comments
			)
			VALUES
			(
				@owner_username,
				@content,
				ARRAY[]::TEXT[]
			)
			RETURNING 
				*
	`

	rows, err := p.server.DB.Pool.Query(ctx, stmt, pgx.NamedArgs{
		"owner_username": owner_username,
		"content":        content,
	})

	if err != nil {
		return nil, fmt.Errorf("Failed to add post: %w", err)
	}

	post, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[model.Post])
	if err != nil {
		return nil, fmt.Errorf("Failed to collect post, %w", err)
	}

	return &post, nil
}

func (p *PostRepository) UpdatePost(ctx context.Context, owner_id int, self_id int, content string) (*model.Post, error) {
	stmt := `
		UPDATE 
			posts 
		SET
				content = @content,
		WHERE 
				owner_id = @owner_id
				AND id = @self_id
		RETURNING 
				*
	`
	rows, err := p.server.DB.Pool.Query(ctx, stmt, pgx.NamedArgs{
		"content":  content,
		"owner_id": owner_id,
		"self_id":  self_id,
	})

	if err != nil {
		return nil, fmt.Errorf("failed to execute update comment query for post_id=%s user_id=%s: %w", self_id, owner_id, err)
	}

	post, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[model.Post])
	if err != nil {
		return nil, fmt.Errorf("failed to collect row from table:posts for post_id=%s user_id=%s: %w", self_id, owner_id, err)
	}
	return &post, nil
}

func (p *PostRepository) GetPost(ctx context.Context, owner_id int, self_id int) (*model.Post, error) {
	stmt := `
		SELECT 
			*
		FROM 
			posts
		WHERE
			owner_id = @owner_id
			AND id = @self_id
	`

	rows, err := p.server.DB.Pool.Query(ctx, stmt, pgx.NamedArgs{
		"owner_id": owner_id,
		"self_id":  self_id,
	})

	if err != nil {
		return nil, fmt.Errorf("failed to execute select post query for post_id=%s user_id=%s: %w", self_id, owner_id, err)
	}

	post, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[model.Post])
	if err != nil {
		return nil, fmt.Errorf("failed to collect row from table:posts for post_id=%s user_id=%s: %w", self_id, owner_id, err)
	}
	return &post, nil
}
