package handler

import (
	"newsapps/internal/features/users"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func ToModelUsers(r RegisterRequest) users.Users {
	return users.Users{
		Name:     r.Name,
		Password: r.Password,
		Email:    r.Email,
	}
}
