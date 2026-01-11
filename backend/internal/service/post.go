package service

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/manish-pandey413/YABS/internal/model"
	"github.com/manish-pandey413/YABS/internal/repository"
	"github.com/manish-pandey413/YABS/internal/server"
)

type PostService struct {
	server   *server.Server
	postRepo *repository.PostRepository
}

func NewPostService(s *server.Server, postRepo *repository.PostRepository) *PostService {
	return &PostService{
		server:   s,
		postRepo: postRepo,
	}
}

func (p *PostService) AddPost(ctx echo.Context, owner_username string, content string) (*model.Post, error) {
	postItem, err := p.postRepo.NewPost(ctx.Request().Context(), owner_username, content)
	if err != nil {
		return nil, fmt.Errorf("Couldn't add post, %w", err)
	}
	return postItem, nil
}
