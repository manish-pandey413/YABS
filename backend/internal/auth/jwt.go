package auth

import (
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/manish-pandey413/YABS/internal/model"
	"github.com/manish-pandey413/YABS/internal/server"
)

func GenJWT(s *server.Server, c echo.Context, user_id int, username string, email string) error {
	claims := &model.Claims{
		User_id:    user_id,
		Username:   username,
		Email:      email,
		Authorized: true,
		Expiration: time.Now().Add(10 * time.Minute).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(s.Config.Auth.SecretKey))
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	tokenCookie := new(http.Cookie)
	tokenCookie.Name = "token"
	tokenCookie.Value = signedToken
	tokenCookie.Expires = time.Now().Add(10 * time.Minute)

	c.SetCookie(tokenCookie)
	return nil
}
