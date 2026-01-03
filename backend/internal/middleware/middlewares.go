package middleware

import "github.com/manish-pandey413/YABS/internal/server"

type Middlewares struct {
	Auth *AuthMiddleware
}

func NewMiddlewares(s *server.Server) *Middlewares {
	return &Middlewares{
		Auth: NewAuthMiddleware(s),
	}
}
