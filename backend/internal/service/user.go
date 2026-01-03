package service

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/manish-pandey413/YABS/internal/model"
	"github.com/manish-pandey413/YABS/internal/repository"
	"github.com/manish-pandey413/YABS/internal/server"
)

type UserService struct {
	server   *server.Server
	userRepo *repository.UserRepository
}

func NewUserService(s *server.Server, userRepo *repository.UserRepository) *UserService {
	return &UserService{
		server:   s,
		userRepo: userRepo,
	}
}

func (u *UserService) AddUser(ctx echo.Context, userName string, email string) (*model.User, error) {
	userItem, err := u.userRepo.AddUser(ctx.Request().Context(), userName, email)
	if err != nil {
		return nil, fmt.Errorf("Couldn't add user with username: %s, %w", userName, err)
	}
	return userItem, nil
}
