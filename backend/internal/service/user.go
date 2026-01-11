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

func (u *UserService) AddUser(ctx echo.Context, userName string, email string, password string) (*model.User, error) {
	userItem, err := u.userRepo.AddUser(ctx.Request().Context(), userName, email, password)
	if err != nil {
		return nil, fmt.Errorf("Couldn't add user, %w", err)
	}
	return userItem, nil
}

func (u *UserService) GetUser(ctx echo.Context, username string) (*model.User, error) {
	userItem, err := u.userRepo.GetUser(ctx.Request().Context(), username)
	if err != nil {
		return nil, fmt.Errorf("Couldn't get user, %w", err)
	}
	return userItem, nil
}
