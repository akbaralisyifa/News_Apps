package utils

import (
	"github.com/go-playground/validator/v10"
)

type AccountUtilityInterface interface {
	EmailPasswordValidator(inputEmail string, inputPw string) error
}

type accountUtility struct {
	vldt validator.Validate
}

// func NewAccountUtility(v validator.Validate) AccountUtilityInterface {
// 	return &accountUtility{
// 		vldt: v,
// 	}
// }

// func (au *accountUtility) EmailPasswordValidator(inputEmail string, inputPw string) error {
// 	err := au.vldt.Struct(&users.LoginValidate{Email: inputEmail, Password: inputPw})
// 	if err != nil {
// 		log.Println("validation error:", err.Error())
// 		return errors.New("validasi gagal")
// 	}
// 	return nil
// }
