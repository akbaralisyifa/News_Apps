package repository

import (
	"newsapps/internal/features/users"

	articleRepository "newsapps/internal/features/articles/repository"
	"newsapps/internal/features/comments/repository"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Password string
	Email    string
	Article  []articleRepository.Articles `gorm:"foreignKey:UserID"`
	Comment  []repository.Comments        `gorm:"foreignKey:UserID"`
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
