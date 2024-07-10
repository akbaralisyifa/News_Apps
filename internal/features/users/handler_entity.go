package users

import (
	"github.com/labstack/echo/v4"
)

type Users struct {
	ID       uint
	Name     string
	Password string
	Email    string
}

type Handler interface {
	Register() echo.HandlerFunc
	Login() echo.HandlerFunc
	UpdateUserAccount() echo.HandlerFunc
	DeleteUserAccount() echo.HandlerFunc
}

type Services interface {
	Register(newUser Users) error
	Login(email string, password string) (Users, string, error)
	UpdateUserAccount(newData Users) error
	DeleteUserAccount(userID uint) error
}

type Query interface {
	Register(newUser Users) error
	Login(email string) (Users, error)
	UpdateUserAccount(newUser Users) error
	DeleteUserAccount(userID uint) error
}

type LoginValidate struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8,alphanum"`
}
type RegisterValidate struct {
	Name     string `validate:"required,min=5,alphanum"`
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=7,alphanum"`
}
