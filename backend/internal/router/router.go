package router

import (
	"github.com/labstack/echo/v4"
	"github.com/manish-pandey413/YABS/internal/handler"
	v1 "github.com/manish-pandey413/YABS/internal/router/v1"
	"github.com/manish-pandey413/YABS/internal/server"
)

func NewRouter(s *server.Server, h *handler.Handlers) *echo.Echo {
	router := echo.New()
	v1Router := router.Group("/api/v1")

	v1.RegisterV1Routes(v1Router, h)

	return router
}
