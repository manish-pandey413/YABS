package handler

import (
	"github.com/manish-pandey413/YABS/internal/server"
	"github.com/manish-pandey413/YABS/internal/service"
)

type Handlers struct {
	UserHandler *UserHandler
	PostHandler *PostHandler
}

func NewHandlers(s *server.Server, services *service.Services) *Handlers {
	return &Handlers{
		UserHandler: NewUserHandler(s, services.User),
		PostHandler: NewPostHandler(s, services.Post),
	}
}
