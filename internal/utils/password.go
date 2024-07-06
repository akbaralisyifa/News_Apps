package utils

import "golang.org/x/crypto/bcrypt"

type PasswordUtilityInterface interface {
	GeneratePassword(currentPw string) ([]byte, error)
	CheckPassword(inputPw []byte, currentPw []byte) error
}
type passwordUtility struct{}

func NewPasswordUtility() PasswordUtilityInterface {
	return &passwordUtility{}
}

func (pw *passwordUtility) GeneratePassword(currentPw string) ([]byte, error) {
	result, err := bcrypt.GenerateFromPassword([]byte(currentPw), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (pw *passwordUtility) CheckPassword(inputPw []byte, currentPw []byte) error {
	return bcrypt.CompareHashAndPassword(currentPw, inputPw)
}
