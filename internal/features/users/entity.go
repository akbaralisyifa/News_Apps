package users

import (
	"github.com/labstack/echo/v4"
)

type Users struct {
	ID       uint
	Name     string
	Password string
	Email    string
	Phone    string
}

type Handler interface {
	Register() echo.HandlerFunc
	Login() echo.HandlerFunc
}

type Services interface {
	Register(newUser Users) error
	Login(email string, password string) (Users, string, error)
}

type Query interface {
	Register(newUser Users) error
	Login(email string) (Users, error)
}

type LoginValidate struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8,alphanum"`
}
type RegisterValidate struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=7,alphanum"`
	Name     string `validate:"required,min=5,alphanum"`
}
