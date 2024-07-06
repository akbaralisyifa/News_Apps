package repository

import (
	"newsapps/internal/features/users"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uint
	Name     string
	Password string
	Email    string
	// Todos     []Todo    `gorm:"foreignKey:Owner"`
}

func (u *User) toUserEntity() users.Users {
	return users.Users{
		ID:       u.ID,
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
	}
}

func toUserQuery(usersfromInputData users.Users) User {
	return User{
		Name:     usersfromInputData.Name,
		Email:    usersfromInputData.Email,
		Password: usersfromInputData.Password,
	}
}
