package handler

import (
	"log"
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
			log.Fatal("Error", err.Error())
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "bad request", nil))
		}
		err = uc.serv.Register(ToModelUsers(input))
		if err != nil {
			log.Fatal("Error", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "server error", nil))
		}
		return c.JSON(http.StatusCreated, helper.ResponseFormat(http.StatusCreated, "success insert data", nil))
	}
}

func (uc *UserController) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input LoginRequest
		err := c.Bind(&input)
		if err != nil {
			log.Fatal("Error", err.Error())
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "bad request", nil))
		}
		result, token, err := uc.serv.Login(input.Email, input.Password)

		if err != nil {
			log.Fatal("Error", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "server error", nil))
		}

		return c.JSON(http.StatusOK, map[string]any{"code": http.StatusOK, "message": "success", "data": ToLoginReponse(result), "token": token})

	}
}
