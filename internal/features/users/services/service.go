package services

import (
	"errors"
	"log"
	"newsapps/internal/features/users"
	"newsapps/internal/utils"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type userServices struct {
	qry  users.Query
	pu   utils.PasswordUtilityInterface
	jwt  utils.JwtUtilityInterface
	vldt utils.AccountUtilityInterface
}

func NewUserService(q users.Query, v utils.AccountUtilityInterface, p utils.PasswordUtilityInterface, j utils.JwtUtilityInterface) users.Services {
	return &userServices{
		qry:  q,
		pu:   p,
		jwt:  j,
		vldt: v,
	}
}

func (us *userServices) Register(newData users.Users) error {

	err := us.vldt.RegisterValidator(newData.Name, newData.Email, newData.Password)
	if err != nil {
		log.Println("login validation error", err.Error())
		return errors.New(err.Error())
	}

	processPw, err := us.pu.GeneratePassword(newData.Password)
	if err != nil {
		log.Println("register generate password error:", err.Error())
		if err.Error() == bcrypt.ErrMismatchedHashAndPassword.Error() {
			return errors.New("data tidak boleh kosong")
		}
		return err
	}
	newData.Password = string(processPw)

	err = us.qry.Register(newData)

	if err != nil {
		log.Println("register sql error:", err.Error())
		return errors.New("terjadi kesalahan pada server saat mengolah data")
	}
	return nil
}

func (us *userServices) Login(email string, password string) (users.Users, string, error) {

	err := us.vldt.EmailPasswordValidator(email, password)
	// Jika validasi gagal
	if err != nil {
		log.Println("validation error:", err.Error())
		return users.Users{}, "", err
	}

	result, err := us.qry.Login(email)
	if err != nil {
		// log.Fatal("Error On Query ", err)
		return users.Users{}, "", gorm.ErrInvalidData
	}

	err = us.pu.CheckPassword([]byte(password), []byte(result.Password))

	if err != nil {
		// log.Fatal("Error On Password", err)
		return users.Users{}, "", errors.New(bcrypt.ErrMismatchedHashAndPassword.Error())
	}

	token, err := us.jwt.GenerateJWT(result.ID, result.Email)
	if err != nil {
		// log.Fatal("Error On Jwt ", err)
		return users.Users{}, "", errors.New("Tidak dapat mendapatkan token")
	}

	return result, token, nil
}
