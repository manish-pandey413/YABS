package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/manish-pandey413/YABS/internal/handler"
)

func registerUserRoutes(r *echo.Group, h *handler.UserHandler) {
	users := r.Group("/users")

	users.POST("/signup", h.Signup)
	users.POST("/login", h.Login)
}
