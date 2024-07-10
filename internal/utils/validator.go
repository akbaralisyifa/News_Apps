package utils

import (
	"errors"
	"log"
	"newsapps/internal/features/articles"
	"newsapps/internal/features/users"

	"github.com/go-playground/validator/v10"
)

type AccountUtilityInterface interface {
	EmailPasswordValidator(inputEmail string, inputPw string) error
	RegisterValidator(inputName string, inputEmail string, inputPw string) error
	CreateArticlesValidator(title string, content string) error 
}

type accountUtility struct {
	vldt validator.Validate
}

func NewAccountUtility(v validator.Validate) AccountUtilityInterface {
	return &accountUtility{
		vldt: v,
	}
}

func (au *accountUtility) EmailPasswordValidator(inputEmail string, inputPw string) error {
	err := au.vldt.Struct(&users.LoginValidate{Email: inputEmail, Password: inputPw})
	if err != nil {
		log.Println("validation error:", err.Error())
		return errors.New("validasi gagal")
	}
	return nil
}

func (au *accountUtility) RegisterValidator(inputName string, inputEmail string, inputPw string) error {
	err := au.vldt.Struct(&users.RegisterValidate{Name: inputName, Email: inputEmail, Password: inputPw})
	if err != nil {
		log.Println("validation error:", err.Error())
		return errors.New("validasi gagal")
	}
	return nil
}

func (au *accountUtility) CreateArticlesValidator(title string, content string) error {
	err := au.vldt.Struct(&articles.ArticlesValidate{Title: title, Content: content});
	if err != nil {
		log.Println("validation error: ", err.Error());
		return errors.New("validation failed")
	};

	return nil;
}