package repository

import (
	"newsapps/internal/features/articles"
	"newsapps/internal/features/comments/repository"

	"gorm.io/gorm"
)

type Articles struct {
	gorm.Model
	UserID   uint
	Title    string
	Content  string
	Image    string
	Comments []repository.Comments `gorm:"foreignKey:article_id"`
}

func (a *Articles) ToArticlesEntity() articles.Article {
	return articles.Article{
		ID:      a.ID,
		Title:   a.Title,
		Content: a.Content,
		Image:   a.Image,
	}
}

func ToArticlesQuery(input articles.Article) Articles {
	return Articles{
		UserID:  input.UserID,
		Title:   input.Title,
		Content: input.Content,
		Image:   input.Image,
	}
}
