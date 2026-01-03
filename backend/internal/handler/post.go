package handler

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/manish-pandey413/YABS/internal/server"
	"github.com/manish-pandey413/YABS/internal/service"
)

type PostHandler struct {
	server      *server.Server
	postService *service.PostService
}

func NewPostHandler(s *server.Server, postService *service.PostService) *PostHandler {
	return &PostHandler{
		server:      s,
		postService: postService,
	}
}

func (p *PostHandler) AddPost(c echo.Context) error {
	fmt.Println("/post route working !")
	return nil
}
