package handler

import "newsapps/internal/features/users"

type LoginResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func ToLoginReponse(input users.Users) LoginResponse {
	return LoginResponse{
		ID:    input.ID,
		Name:  input.Name,
		Email: input.Email,
	}
}
