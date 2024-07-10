package handler

import "newsapps/internal/features/articles"

type ArticlesRequeste struct {
	UserID  uint   `json:"user_id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Image   string `json:"image"`
}

func ToRequeteArticles(ar ArticlesRequeste) articles.Article {
	return articles.Article{
		UserID:  ar.UserID,
		Title:   ar.Title,
		Content: ar.Content,
		Image:   ar.Image,
	}
}
