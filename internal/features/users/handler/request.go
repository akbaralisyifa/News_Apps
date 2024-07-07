package handler

import (
	"newsapps/internal/features/users"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type UpdateAccountRequest struct {
	UserID   uint   `json:"user_id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func ToModelUsers(r UserRequest) users.Users {
	return users.Users{
		Name:     r.Name,
		Password: r.Password,
		Email:    r.Email,
	}
}

func ToModelUsersAccount(r UpdateAccountRequest) users.Users {
	return users.Users{
		ID:       r.UserID,
		Name:     r.Name,
		Password: r.Password,
		Email:    r.Email,
	}
}
