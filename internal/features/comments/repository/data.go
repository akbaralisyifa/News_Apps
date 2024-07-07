package repository

import "gorm.io/gorm"

type Comments struct {
	gorm.Model
	UserID    uint
	ArticleID uint
	Comment   string
}
