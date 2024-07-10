package articles

import "github.com/labstack/echo/v4"

type Article struct {
	ID      uint
	UserID  uint
	Title   string
	Content string
	Image   string
	Comments []Comment
}

type Comment struct {
	UserID  uint
	Comment string
}

type Handler interface {
	CreateArticles() echo.HandlerFunc
	GetArticles() echo.HandlerFunc
	UpdateArticles() echo.HandlerFunc
	DeleteArticles() echo.HandlerFunc
	GetArticlesByID() echo.HandlerFunc
}

type Services interface {
	CreateArticles(newArticles Article) error
	GetArticles() ([]Article, error)
	UpdateArticles(id uint, updateArticles Article) error
	DeleteArticles(id uint, userID uint) error
	GetArticlesByID(id uint)(Article, error)
}

type Query interface {
	GetArticles() ([]Article, error)
	CreateArticles(newArticles Article) error
	UpdateArticles(id uint, updateArticles Article) error
	DeleteArticles(id uint, userID uint) error
	GetArticlesByID(id uint)(Article, error)
}

// initial validator
type ArticlesValidate struct {
	Title   string `validate:"required"`
	Content string	`validate:"required"`
}