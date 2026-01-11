package middleware

import (
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt"

	"github.com/labstack/echo/v4"
	"github.com/manish-pandey413/YABS/internal/model"
	"github.com/manish-pandey413/YABS/internal/server"
)

type AuthMiddleware struct {
	server *server.Server
}

func NewAuthMiddleware(s *server.Server) *AuthMiddleware {
	return &AuthMiddleware{
		server: s,
	}
}

func (a *AuthMiddleware) RequireAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		cookie, err := c.Cookie("token")
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"error": "Couldn't get cookie" + err.Error(),
			})
		}

		// tokenFromHeader := c.Request().Header.Get(echo.HeaderAuthorization)
		tokenFromHeader := cookie.Value
		if tokenFromHeader == "" {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "No Token provided"})
		}

		claims := &model.Claims{}

		parsedToken, err := jwt.ParseWithClaims(tokenFromHeader, claims, func(token *jwt.Token) (any, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}
			return []byte(a.server.Config.Auth.SecretKey), nil
		})

		if !parsedToken.Valid || err != nil {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Invalid token"})
		}

		c.Set("owner_username", claims.Username)

		return next(c)
	}
}
