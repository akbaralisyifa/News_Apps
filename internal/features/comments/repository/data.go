package repository

import (
	"newsapps/internal/features/comments"

	"gorm.io/gorm"
)

type Comments struct {
	gorm.Model
	UserID    uint
	ArticleID uint
	Comment   string
}

func (a *Comments) ToCommentsEntity() comments.Comment {
	return comments.Comment{
		ID:        a.ID,
		UserID:    a.UserID,
		ArticleID: a.ArticleID,
		Comments:  a.Comment,
	}
}

func ToArticlesQuery(input comments.Comment) Comments {
	return Comments{

		UserID:    input.UserID,
		ArticleID: input.ArticleID,
		Comment:   input.Comments,
	}
}
