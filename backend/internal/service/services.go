package service

import (
	"github.com/manish-pandey413/YABS/internal/repository"
	"github.com/manish-pandey413/YABS/internal/server"
)

type Services struct {
	User *UserService
	Post *PostService
}

func NewServices(s *server.Server, repos *repository.Repositories) (*Services, error) {
	return &Services{
		User: NewUserService(s, repos.User),
		Post: NewPostService(s, repos.Post),
	}, nil
}
