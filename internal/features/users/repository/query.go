package repository

import (
	"newsapps/internal/features/users"

	"gorm.io/gorm"
)

type UserModel struct {
	db *gorm.DB
}

func NewUserModel(connection *gorm.DB) users.Query {
	return &UserModel{
		db: connection,
	}
}

func (um *UserModel) Login(email string) (users.Users, error) {
	var result User
	err := um.db.Where("email = ?", email).First(&result).Error
	if err != nil {
		return users.Users{}, err
	}
	return result.toUserEntity(), nil
}

func (um *UserModel) Register(newUser users.Users) error {
	resultData := toUserQuery(newUser)
	err := um.db.Create(&resultData).Error
	return err
}