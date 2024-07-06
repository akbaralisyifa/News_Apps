package handler

import (
	"newsapps/internal/features/articles"
)

type ArticlesController struct {
	srv articles.Services
}

func NewArticlesController(s articles.Services) articles.Handler {
	return &ArticlesController{
		srv : s,
	}
}
