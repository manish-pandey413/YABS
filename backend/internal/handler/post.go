package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/manish-pandey413/YABS/internal/model"
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
	owner_username := c.Get("owner_username").(string)
	recievedPost := &model.Post{}
	if err := c.Bind(recievedPost); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "Invalid request format: " + err.Error()})
	}

	postItem, err := p.postService.AddPost(c, owner_username, recievedPost.Content)
	if err != nil {
		fmt.Printf("Couldn't add post to username:-  %s :%s", owner_username, err)
		return fmt.Errorf("%w", err)
	}
	return c.JSON(http.StatusCreated, postItem)
}
