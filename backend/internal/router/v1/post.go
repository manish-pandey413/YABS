package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/manish-pandey413/YABS/internal/handler"
	"github.com/manish-pandey413/YABS/internal/middleware"
)

func registerPostRoutes(r *echo.Group, h *handler.PostHandler, m *middleware.AuthMiddleware) {
	posts := r.Group("/posts")
	posts.Use(m.RequireAuth)

	posts.POST("", h.AddPost)
}
