package router

import (
	"github.com/labstack/echo/v4"
	"github.com/manish-pandey413/YABS/internal/handler"
	"github.com/manish-pandey413/YABS/internal/middleware"
	v1 "github.com/manish-pandey413/YABS/internal/router/v1"
	"github.com/manish-pandey413/YABS/internal/server"
)

func NewRouter(s *server.Server, h *handler.Handlers) *echo.Echo {
	middlewares := middleware.NewMiddlewares(s)
	router := echo.New()
	v1Router := router.Group("/api/v1")

	v1.RegisterV1Routes(v1Router, h, middlewares)

	return router
}
