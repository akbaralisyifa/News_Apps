package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"newsapps/configs"
)

type JwtUtilityInterface interface {
	GenerateJWT(id uint, email string) (string, error)
	DecodeToken(token *jwt.Token) float64
}
type jwtUtility struct{}

func NewJwtUtility() JwtUtilityInterface {
	return &jwtUtility{}
}

func (j *jwtUtility) GenerateJWT(id uint, email string) (string, error) {

	passKey := configs.ImportSetting().JWTSECRET
	var data = jwt.MapClaims{}
	// custom data
	data["id"] = id
	data["email"] = email
	// mandatory data
	data["iat"] = time.Now().Unix()
	data["exp"] = time.Now().Add(time.Hour * 3).Unix()

	var proccessToken = jwt.NewWithClaims(jwt.SigningMethodHS256, data)

	result, err := proccessToken.SignedString([]byte(passKey))

	if err != nil {
		return "", err
	}

	return result, nil
}

func (j *jwtUtility) DecodeToken(token *jwt.Token) float64 {
	var result float64
	var claim = token.Claims.(jwt.MapClaims)

	for _, v := range claim {
		fmt.Println(v)
	}

	if val, found := claim["id"]; found {
		result = val.(float64)
	}

	return result
}
