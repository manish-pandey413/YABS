package handler

import (
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/labstack/echo/v4"
	"github.com/manish-pandey413/YABS/internal/auth"
	"github.com/manish-pandey413/YABS/internal/model"
	"github.com/manish-pandey413/YABS/internal/server"
	"github.com/manish-pandey413/YABS/internal/service"
)

type UserHandler struct {
	server      *server.Server
	userService *service.UserService
}

func NewUserHandler(s *server.Server, userService *service.UserService) *UserHandler {
	return &UserHandler{
		server:      s,
		userService: userService,
	}
}

func (u *UserHandler) Signup(c echo.Context) error {
	userCred := &model.User{}
	userCred.User_id = 0

	if err := c.Bind(userCred); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "Invalid request format: " + err.Error()})
	}

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(userCred.Password), 10)
	if err != nil {
		fmt.Println("Couldn't hash password: ", err)
		return fmt.Errorf("%w", err)
	}

	userItem, err := u.userService.AddUser(c, userCred.Username, userCred.Email, string(hashedPass))
	if err != nil {
		fmt.Println("Couldn't add user: ", err)
		return fmt.Errorf("%w", err)
	}

	return c.JSON(http.StatusCreated, userItem)
}

func (u *UserHandler) Login(c echo.Context) error {
	recievedCreds := &model.UserCred{}
	if err := c.Bind(recievedCreds); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "Invalid request format: " + err.Error()})
	}

	user, err := u.userService.GetUser(c, recievedCreds.Username)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "Invalid username or password"})
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(recievedCreds.Password)); err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"error": "Invalid username or password",
		})
	}

	if err := auth.GenJWT(u.server, c, user.User_id, user.Username, user.Email); err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"error": "Couldn't generate JWT" + err.Error(),
		})
	}

	return c.JSON(http.StatusAccepted, echo.Map{"message": "Welcome!"})
}
