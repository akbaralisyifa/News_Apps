package handler

import (
	"net/http"
	"newsapps/internal/features/users"
	"newsapps/internal/helper"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	serv users.Services
}

func NewUserController(s users.Services) users.Handler {
	return &UserController{
		serv: s,
	}
}

func (uc *UserController) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input RegisterRequest
		err := c.Bind(&input)
		if err != nil {
			return c.JSON(400, helper.ResponseFormat(400, "bad request", nil))
		}
		err = uc.serv.Register(ToModelUsers(input))
		if err != nil {
			return c.JSON(500, helper.ResponseFormat(500, "server error", nil))
		}
		return c.JSON(201, helper.ResponseFormat(201, "success insert data", nil))
	}
}

func (uc *UserController) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input LoginRequest
		err := c.Bind(&input)
		if err != nil {
			return c.JSON(400, helper.ResponseFormat(400, "bad request", nil))
		}
		result, token, err := uc.serv.Login(input.Email, input.Password)

		if err != nil {
			return c.JSON(500, helper.ResponseFormat(500, "server error", nil))
		}

		return c.JSON(http.StatusOK, map[string]any{"code": http.StatusOK, "message": "selama anda berhasil login", "data": ToLoginReponse(result), "token": token})

	}
}
