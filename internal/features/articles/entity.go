package articles

import "github.com/labstack/echo/v4"

type Article struct {
	ID      uint
	UserID  uint
	Title   string
	Content string
	Image   string
}

type Handler interface {
	CreateArticles() echo.HandlerFunc
	GetArticles() echo.HandlerFunc
	UpdateArticles() echo.HandlerFunc
	DeleteArticles() echo.HandlerFunc
}

type Services interface {
	CreateArticles(newArticles Article) error
	GetArticles() ([]Article, error)
	UpdateArticles(id uint, updateArticles Article) error
	DeleteArticles(id uint) error
	GetArticlesByID(id uint)(Article, error)
}

type Query interface {
	GetArticles() ([]Article, error)
	CreateArticles(newArticles Article) error
	UpdateArticles(id uint, updateArticles Article) error
	DeleteArticles(id uint) error
	GetArticlesByID(id uint)(Article, error)
}