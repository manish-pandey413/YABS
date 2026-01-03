package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/manish-pandey413/YABS/internal/handler"
)

func RegisterV1Routes(router *echo.Group, handlers *handler.Handlers) {
	registerUserRoutes(router, handlers.UserHandler)
	registerPostRoutes(router, handlers.PostHandler)
}
