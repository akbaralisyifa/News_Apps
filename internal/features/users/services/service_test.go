package services_test

import (
	"errors"
	"newsapps/internal/features/users"
	"newsapps/internal/features/users/services"
	"newsapps/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func TestRegister(t *testing.T) {
	qry := mocks.NewQuery(t)
	pu := mocks.NewPasswordUtilityInterface(t)
	jwt := mocks.NewJwtUtilityInterface(t)
	vldt := mocks.NewAccountUtilityInterface(t)
	srv := services.NewUserService(qry, vldt, pu, jwt)

	input := users.Users{Name: "anggieko", Password: "anggi12345", Email: "anggi@eko.com"}

	t.Run("Success Register", func(t *testing.T) {
		inputQry := users.Users{Name: "anggieko", Password: "somepassword", Email: "anggi@eko.com"}

		vldt.On("RegisterValidator", input.Name, input.Email, input.Password).Return(nil).Once()
		pu.On("GeneratePassword", input.Password).Return([]byte("somepassword"), nil).Once()
		qry.On("Register", inputQry).Return(nil).Once()

		err := srv.Register(input)

		vldt.AssertExpectations(t)
		pu.AssertExpectations(t)
		qry.AssertExpectations(t)

		assert.Nil(t, err)
		// assert.Equal(t, "anggi12345", input.Password)

	})

	t.Run("Error Hash Password", func(t *testing.T) {
		originalPassword := input.Password
		vldt.On("RegisterValidator", input.Name, input.Email, input.Password).Return(nil).Once()
		pu.On("GeneratePassword", input.Password).Return(nil, bcrypt.ErrPasswordTooLong).Once()

		err := srv.Register(input)

		vldt.AssertExpectations(t)
		pu.AssertExpectations(t)

		assert.Error(t, err)
		assert.ErrorContains(t, err, bcrypt.ErrPasswordTooLong.Error())
		assert.Equal(t, originalPassword, input.Password) // memastikan password tetap sama

	})

	t.Run("Error From Query", func(t *testing.T) {
		inputQry := users.Users{Name: "anggieko", Password: "goodpassword", Email: "anggi@eko.com"}
		vldt.On("RegisterValidator", input.Name, input.Email, input.Password).Return(nil).Once()
		pu.On("GeneratePassword", input.Password).Return([]byte("goodpassword"), nil).Once()
		qry.On("Register", inputQry).Return(gorm.ErrInvalidData).Once()

		err := srv.Register(input)

		vldt.AssertExpectations(t)
		pu.AssertExpectations(t)
		qry.AssertExpectations(t)

		assert.Error(t, err)
		assert.ErrorContains(t, err, "terjadi kesalahan pada server saat mengolah data")

	})

	t.Run("Nil Value", func(t *testing.T) {
		vldt.On("RegisterValidator", input.Name, input.Email, input.Password).Return(errors.New("validasi tidak sesuai")).Once()

		err := srv.Register(input)

		vldt.AssertExpectations(t)

		assert.Error(t, err)
		assert.ErrorContains(t, err, "validasi tidak sesuai")
	})

}

func TestLogin(t *testing.T) {
	qry := mocks.NewQuery(t)
	pu := mocks.NewPasswordUtilityInterface(t)
	jwt := mocks.NewJwtUtilityInterface(t)
	vldt := mocks.NewAccountUtilityInterface(t)
	srv := services.NewUserService(qry, vldt, pu, jwt)

	input := users.Users{Password: "anggi1234", Email: "anggi@eko.com"}

	t.Run("Success Login", func(t *testing.T) {
		inputQry := users.Users{Password: "anggi1234", Email: "anggi@eko.com"}

		vldt.On("EmailPasswordValidator", input.Email, input.Password).Return(nil).Once()
		qry.On("Login", input.Email).Return(inputQry, nil).Once()
		pu.On("CheckPassword", []byte(inputQry.Password), []byte(inputQry.Password)).Return(nil).Once()
		jwt.On("GenerateJWT", input.ID, input.Email).Return("someToken", nil).Once()

		user, token, err := srv.Login(inputQry.Email, inputQry.Password)

		vldt.AssertExpectations(t)
		qry.AssertExpectations(t)
		pu.AssertExpectations(t)
		jwt.AssertExpectations(t)

		assert.Nil(t, err)
		assert.Equal(t, "someToken", token)
		assert.Equal(t, inputQry, user)
	})

	t.Run("Failed email/password Login", func(t *testing.T) {
		inputQry := users.Users{Password: "", Email: ""}

		vldt.On("EmailPasswordValidator", inputQry.Email, inputQry.Password).Return(errors.New("validasi gagal")).Once()

		_, _, err := srv.Login(inputQry.Email, inputQry.Password)

		vldt.AssertExpectations(t)

		assert.Error(t, err)
		assert.ErrorContains(t, err, "validasi gagal")
	})

	t.Run("Error on Query", func(t *testing.T) {
		inputQry := users.Users{Password: "anggi1234", Email: "anggi@eko.com"}

		vldt.On("EmailPasswordValidator", inputQry.Email, inputQry.Password).Return(nil).Once()
		qry.On("Login", input.Email).Return(inputQry, gorm.ErrInvalidData).Once()
		// pu.On("CheckPassword", []byte(inputQry.Password), []byte(inputQry.Password)).Return(nil).Once()

		_, _, err := srv.Login(inputQry.Email, inputQry.Password)

		vldt.AssertExpectations(t)
		qry.AssertExpectations(t)

		assert.Error(t, err)
		assert.ErrorContains(t, err, string(gorm.ErrInvalidData.Error()))
	})

	t.Run("Error on Password", func(t *testing.T) {
		inputQry := users.Users{Password: "hashedPassword", Email: "anggi@eko.com"}

		vldt.On("EmailPasswordValidator", inputQry.Email, inputQry.Password).Return(nil).Once()
		qry.On("Login", input.Email).Return(inputQry, nil).Once()
		pu.On("CheckPassword", []byte(inputQry.Password), []byte(inputQry.Password)).Return(bcrypt.ErrMismatchedHashAndPassword).Once()

		_, _, err := srv.Login(inputQry.Email, inputQry.Password)

		vldt.AssertExpectations(t)
		qry.AssertExpectations(t)
		pu.AssertExpectations(t)

		assert.Error(t, err)
		assert.ErrorContains(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())
	})

	t.Run("Error on JWT Generator", func(t *testing.T) {
		inputQry := users.Users{Password: "hashedPassword", Email: "anggi@eko.com"}

		vldt.On("EmailPasswordValidator", inputQry.Email, inputQry.Password).Return(nil).Once()
		qry.On("Login", input.Email).Return(inputQry, nil).Once()
		pu.On("CheckPassword", []byte(inputQry.Password), []byte(inputQry.Password)).Return(nil).Once()
		jwt.On("GenerateJWT", input.ID, input.Email).Return("", errors.New("Tidak dapat mendapatkan token")).Once()

		_, _, err := srv.Login(inputQry.Email, inputQry.Password)

		vldt.AssertExpectations(t)
		qry.AssertExpectations(t)
		pu.AssertExpectations(t)
		jwt.AssertExpectations(t)

		assert.Error(t, err)
		assert.ErrorContains(t, err, "Tidak dapat mendapatkan token")
	})
}
