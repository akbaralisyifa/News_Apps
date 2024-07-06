package repository

import (
	"newsapps/internal/features/comments/repository"

	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	ID      uint
	Title   string
	Content string
	User_id uint
	Image   string
	Comment []repository.Comment `gorm:"foreignKey:Article_id"`
}
