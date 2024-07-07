package handler

import (
	"log"
	"net/http"
	"newsapps/internal/features/users"
	"newsapps/internal/helper"
	"newsapps/internal/utils"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	srv users.Services
}

func NewUserController(s users.Services) users.Handler {
	return &UserController{
		srv: s,
	}
}

func (uc *UserController) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input UserRequest
		err := c.Bind(&input)
		if err != nil {
			log.Fatal("Error", err.Error())
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "bad request", nil))
		}
		err = uc.srv.Register(ToModelUsers(input))
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
		result, token, err := uc.srv.Login(input.Email, input.Password)

		if err != nil {
			log.Fatal("Error", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "server error", nil))
		}

		return c.JSON(http.StatusOK, map[string]any{"code": http.StatusOK, "message": "success", "data": ToLoginReponse(result), "token": token})

	}
}

func (uc *UserController) UpdateUserAccount() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input UpdateAccountRequest
		err := c.Bind(&input)
		if err != nil {
			c.Logger().Error("error bind input: ", err.Error())
			return c.JSON(400, helper.ResponseFormat(400, "input error", nil))
		}

		var userIDJWT = utils.NewJwtUtility().DecodeToken(c.Get("user").(*jwt.Token))

		if input.UserID != uint(userIDJWT) {
			c.Logger().Error("id not matched")
			return c.JSON(401, helper.ResponseFormat(401, "Unauthorized", nil))
		}

		err = uc.srv.UpdateUserAccount(ToModelUsersAccount(input))
		if err != nil {
			c.Logger().Error("when update data error: ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "server error", nil))
		}
		return c.JSON(http.StatusCreated, helper.ResponseFormat(http.StatusCreated, "success insert data", nil))
	}
}

func (uc *UserController) DeleteUserAccount() echo.HandlerFunc {
	return func(c echo.Context) error {

		// id, _ := strconv.Atoi(c.Param("id"))
		var userID float64 = 0
		userID = utils.NewJwtUtility().DecodeToken(c.Get("user").(*jwt.Token))
		if userID == 0 {
			return c.JSON(500, helper.ResponseFormat(400, "Invalid JWT", nil))
		}
		err := uc.srv.DeleteUserAccount(uint(userID))

		if err != nil {
			return c.JSON(500, helper.ResponseFormat(500, "server error", nil))
		}

		return c.JSON(200, helper.ResponseFormat(200, "success delete data", nil))
	}
}
