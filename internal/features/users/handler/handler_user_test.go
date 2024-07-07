package handler_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"newsapps/internal/features/users"
	"newsapps/internal/features/users/handler"
	"newsapps/internal/helper"
	"newsapps/mocks"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRegister(t *testing.T) {
	srv := mocks.NewServices(t)
	userController := handler.NewUserController(srv)

	t.Run("Success Register", func(t *testing.T) {
		srv.On("Register", mock.Anything).Return(nil).Once()

		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/register", strings.NewReader(`{"email":"test@example.com", "password":"password123"}`))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err := userController.Register()(c)
		if assert.NoError(t, err) {
			assert.Equal(t, http.StatusCreated, rec.Code)
			expectedResponse := helper.ResponseFormat(http.StatusCreated, "success insert data", nil)
			jsonBytes, _ := json.Marshal(expectedResponse)

			assert.JSONEq(t, string(jsonBytes), rec.Body.String())
		}

		srv.AssertExpectations(t)
	})

	t.Run("Bad Request - Invalid Json", func(t *testing.T) {

		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/register", strings.NewReader(`invalid json`))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err := userController.Register()(c)
		if assert.NoError(t, err) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			expectedResponse := helper.ResponseFormat(http.StatusBadRequest, "bad request", nil)
			jsonBytes, _ := json.Marshal(expectedResponse)
			assert.JSONEq(t, string(jsonBytes), rec.Body.String())
		}

		srv.AssertExpectations(t)
	})

	t.Run("Server Error", func(t *testing.T) {
		srv.On("Register", mock.Anything).Return(errors.New("internal server error")).Once()

		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/register", strings.NewReader(`{"email":"test@example.com", "password":"password123"}`))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err := userController.Register()(c)
		if assert.NoError(t, err) {
			assert.Equal(t, http.StatusInternalServerError, rec.Code)
			expectedResponse := helper.ResponseFormat(http.StatusInternalServerError, "server error", nil)
			jsonBytes, _ := json.Marshal(expectedResponse)

			assert.JSONEq(t, string(jsonBytes), rec.Body.String())
		}

		srv.AssertExpectations(t)
	})
}

func TestLogin(t *testing.T) {
	srv := mocks.NewServices(t)
	userController := handler.NewUserController(srv)

	t.Run("Success Login", func(t *testing.T) {
		exampleUsersResponse := users.Users{
			ID:       1,
			Name:     "yourname",
			Email:    "test@example.com",
			Password: "",
		}
		srv.On("Login", "test@example.com", "password123").Return(exampleUsersResponse, "token", nil).Once()
		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(`{"email":"test@example.com", "password":"password123"}`))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err := userController.Login()(c)
		if assert.NoError(t, err) {

			exampleUsersResponse := fmt.Sprintf(`{"code": %d, "data":{
			"id": %d,
			"name": "%s",
			"email": "%s"
			},"message": "success","token": "token"}`,
				http.StatusOK, exampleUsersResponse.ID, exampleUsersResponse.Name, exampleUsersResponse.Email)

			assert.Equal(t, http.StatusOK, rec.Code)
			assert.JSONEq(t, string([]byte(exampleUsersResponse)), rec.Body.String())
		}

		srv.AssertExpectations(t)
	})

	t.Run("Bad Request", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(`invalid json`))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err := userController.Login()(c)
		if assert.NoError(t, err) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			expectedResponse := helper.ResponseFormat(http.StatusBadRequest, "bad request", nil)
			jsonBytes, _ := json.Marshal(expectedResponse)
			assert.JSONEq(t, string(jsonBytes), rec.Body.String())
		}

		srv.AssertExpectations(t)
	})

	t.Run("Server Error", func(t *testing.T) {
		srv.On("Login", "test@example.com", "password123").Return(users.Users{}, "", errors.New("internal server error")).Once()

		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(`{"email":"test@example.com", "password":"password123"}`))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err := userController.Login()(c)
		if assert.NoError(t, err) {
			assert.Equal(t, http.StatusInternalServerError, rec.Code)
			expectedResponse := helper.ResponseFormat(http.StatusInternalServerError, "server error", nil)
			jsonBytes, _ := json.Marshal(expectedResponse)

			assert.JSONEq(t, string(jsonBytes), rec.Body.String())
		}

		srv.AssertExpectations(t)
	})
}
