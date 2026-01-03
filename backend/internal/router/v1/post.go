package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/manish-pandey413/YABS/internal/handler"
)

func registerPostRoutes(r *echo.Group, h *handler.PostHandler) {
	posts := r.Group("/Posts")

	posts.POST("", h.AddPost)
}
