package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/manish-pandey413/YABS/internal/handler"
	"github.com/manish-pandey413/YABS/internal/middleware"
)

func RegisterV1Routes(router *echo.Group, handlers *handler.Handlers, middlewares *middleware.Middlewares) {
	registerUserRoutes(router, handlers.UserHandler)
	registerPostRoutes(router, handlers.PostHandler, middlewares.Auth)
}
