package repository

import (
	"fmt"
	"log"
	"newsapps/internal/features/users"
	"strings"
	"time"

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

func (um *UserModel) UpdateUserAccount(newUser users.Users) error {
	// query := `UPDATE "newsapps"."users" SET AND "name"= ?, "email": ?, password = ?, "updated_at"= ?
	// WHERE ID = ? AND "todos"."deleted_at" IS NULL;`
	// err := um.db.Exec(query, &newUser.Name, &newUser.Email, &newUser.Password, time.Now(), &newUser.ID).Error
	// fmt.Println(query)

	query := `UPDATE "newsapps"."users" SET "updated_at" = $1`
	args := []interface{}{time.Now()}
	index := 2

	setClauses := []string{}
	if newUser.Name != "" {
		setClauses = append(setClauses, fmt.Sprintf(`"name" = $%d`, index))
		args = append(args, newUser.Name)
		index++
	}
	if newUser.Email != "" {
		setClauses = append(setClauses, fmt.Sprintf(`"email" = $%d`, index))
		args = append(args, newUser.Email)
		index++
	}
	if newUser.Password != "" {
		setClauses = append(setClauses, fmt.Sprintf(`"password" = $%d`, index))
		args = append(args, newUser.Password)
		index++
	}

	if len(setClauses) == 0 {
		log.Printf("error len")
		return fmt.Errorf("no valid fields to update")
	}

	query += ", " + strings.Join(setClauses, ", ")
	query += ` WHERE "id" = $` + fmt.Sprintf("%d", index)
	args = append(args, newUser.ID)
	index++
	query += ` AND "users"."deleted_at" IS NULL`

	log.Printf("Executing query: %s with args: %v", query, args)

	err := um.db.Debug().Exec(query, args...)
	if err.Error != nil {
		log.Printf("Query error: %v", err.Error)
		return fmt.Errorf("query error")
	}
	return nil
}

func (um *UserModel) DeleteUserAccount(userID uint) error {
	qry := um.db.Where("id = ?", userID).Delete(&User{})

	if qry.Error != nil {
		log.Printf("Query error: %v", qry.Error)
		return qry.Error
	}

	if qry.RowsAffected < 1 {
		log.Printf("error row affected 0")
		return gorm.ErrRecordNotFound
	}

	return nil
}
