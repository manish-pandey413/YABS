package handler

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/manish-pandey413/YABS/internal/server"
	"github.com/manish-pandey413/YABS/internal/service"
)

type UserHandler struct {
	server      *server.Server
	userService *service.UserService
}

func NewUserHandler(s *server.Server, userService *service.UserService) *UserHandler {
	return &UserHandler{
		server:      s,
		userService: userService,
	}
}

func (h *UserHandler) AddUser(c echo.Context) error {
	fmt.Println("/user route working !")
	return nil
}
