package repository

import "github.com/manish-pandey413/YABS/internal/server"

type Repositories struct {
	User *UserRepository
	Post *PostRepository
}

func NewRepositories(s *server.Server) *Repositories {
	return &Repositories{
		User: NewUserRepository(s),
		Post: NewPostRepository(s),
	}
}
