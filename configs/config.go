package configs

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type setting struct {
	Host      string
	User      string
	Password  string
	Port      string
	DBNAME    string
	JWTSECRET string
}

func ImportSetting() setting {
	var result setting
	err := godotenv.Load(".env")
	if err != nil {
		return setting{}
	}
	result.Host = os.Getenv("poshost")
	result.User = os.Getenv("posuser")
	result.Password = os.Getenv("pospw")
	result.Port = os.Getenv("posport")
	result.DBNAME = os.Getenv("dbname")
	result.JWTSECRET = os.Getenv("JWT_SECRET")

	return result
}

func ConnectDB(s setting) (*gorm.DB, error) {
	var connStr = fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s", s.Host, s.User, s.Password, s.Port, s.DBNAME)
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "be23.",
		},
	})

	if err != nil {
		return nil, err
	}

	return db, nil
}
