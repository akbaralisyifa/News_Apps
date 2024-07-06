package repository

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	ID         uint
	Article_id uint
	User_id    uint
	Comment    string
}
