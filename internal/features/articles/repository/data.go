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
		Comments: nil,
	}
}

func ToArticlesEntityGetAll(articlesList []Articles) []articles.Article{
	articlesEntity := make([]articles.Article, len(articlesList));

		for i, val := range articlesList{
			articlesEntity[i] = val.ToArticlesEntityComments()
		}

	return articlesEntity;
}

func(a *Articles) ToArticlesEntityComments() articles.Article{
	articlesEntity := a.ToArticlesEntity();

	if len(a.Comments) > 0 {
		articlesEntity.Comments = make([]string, len(a.Comments));
		for i, val := range a.Comments{
			articlesEntity.Comments[i] = val.Comment
		}
	}

	return articlesEntity;
}

func ToArticlesQuery(input articles.Article) Articles {
	return Articles{
		UserID:  input.UserID,
		Title:   input.Title,
		Content: input.Content,
		Image:   input.Image,
	}
}
